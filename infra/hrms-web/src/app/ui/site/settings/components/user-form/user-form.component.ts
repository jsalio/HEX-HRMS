import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { CreateUser, UserType } from '../../../../../core/domain/models';

@Component({
  selector: 'app-user-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './user-form.component.html',
  styleUrl: './user-form.component.css'
})
export class UserFormComponent {
  @Input() saving = false;
  @Input() error: string | null = null;
  @Output() save = new EventEmitter<CreateUser>();
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

    // Sync username with email since UI only has one field
    this.form.get('username')?.valueChanges.subscribe(value => {
      this.form.patchValue({ email: value }, { emitEvent: false });
    });
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
