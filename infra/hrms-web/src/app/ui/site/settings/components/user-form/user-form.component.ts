import { Component, EventEmitter, Input, Output, OnChanges, OnInit, SimpleChanges, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { CreateUser, UserData, UserType, Role } from '../../../../../core/domain/models';
import { ListRoleUseCase } from '../../../../../core/usecases/role';

@Component({
  selector: 'app-user-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './user-form.component.html',
  styleUrl: './user-form.component.css'
})
export class UserFormComponent implements OnChanges, OnInit {
  @Input() saving = false;
  @Input() error: string | null = null;
  @Input() user: UserData | null = null;
  @Output() save = new EventEmitter<CreateUser>();
  @Output() cancel = new EventEmitter<void>();

  private listRoleUseCase = inject(ListRoleUseCase);

  form: FormGroup;
  userTypes = Object.values(UserType);
  roles = signal<Role[]>([]);

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

    // Sync username with email
    this.form.get('username')?.valueChanges.subscribe(value => {
      this.form.patchValue({ email: value }, { emitEvent: false });
    });
  }

  ngOnInit(): void {
    this.fetchRoles();
  }

  fetchRoles(): void {
    this.listRoleUseCase.Execute({
      filters: [],
      pagination: { page: 1, limit: 100 }
    }).then(data => {
      this.roles.set(data.rows);
    }).catch(err => {
      console.error('Error fetching roles:', err);
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
      
      // Make password optional in edit mode
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
