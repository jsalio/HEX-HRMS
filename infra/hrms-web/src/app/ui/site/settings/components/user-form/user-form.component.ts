import { Component, EventEmitter, Input, Output, OnChanges, SimpleChanges } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { CreateUser, UserData, UserType } from '../../../../../core/domain/models';

@Component({
  selector: 'app-user-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './user-form.component.html',
  styleUrl: './user-form.component.css'
})
export class UserFormComponent implements OnChanges {
  @Input() saving = false;
  @Input() error: string | null = null;
  @Input() user: UserData | null = null;
  @Output() save = new EventEmitter<CreateUser>(); // Or ModifyUser but structure is similar enough often
  @Output() cancel = new EventEmitter<void>();

  form: FormGroup;
  userTypes = Object.values(UserType);

  constructor(private fb: FormBuilder) {
    this.form = this.fb.group({
      username: ['', [Validators.required, Validators.email]],
      name: ['', Validators.required],
      lastName: ['', Validators.required],
      password: ['', [Validators.required, Validators.minLength(6)]],
      email: ['', [Validators.required, Validators.email]],
      type: [UserType.Normal, Validators.required],
      role: ['', Validators.required],
      picture: ['']
    });

    // Sync username with email since UI only has one field (if creating)
    this.form.get('username')?.valueChanges.subscribe(value => {
       // Only sync if not editing? Or always? Assuming always for now.
      this.form.patchValue({ email: value }, { emitEvent: false });
    });
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes['user'] && this.user) {
      this.form.patchValue({
        username: this.user.username,
        name: this.user.name,
        lastName: this.user.lastName,
        email: this.user.email,
        type: this.user.type,
        role: this.user.role,
        picture: this.user.picture
      });
      // Handle password - usually empty on edit?
      // For now, if editing, we might need to remove password validator or keep it if password change is forced/allowed.
      // If user is passed, it's edit mode. Make password optional if not changing?
      // Simpler: Keep it as is. User needs to re-enter password to save? No, that's bad UX.
      // Better: Remove required validator from password if in edit mode.
      
      this.form.get('password')?.clearValidators();
      this.form.get('password')?.updateValueAndValidity();
    }
  }

  onSubmit() {
    if (this.form.valid) {
      this.save.emit(this.form.value);
    } else {
      console.log('Form invalid:', this.form.errors);
      Object.keys(this.form.controls).forEach(key => {
        const control = this.form.get(key);
        if (control?.invalid) {
          console.log(`Field ${key} is invalid:`, control.errors);
          control.markAsTouched();
        }
      });
    }
  }

  onCancel() {
    this.cancel.emit();
  }
}
