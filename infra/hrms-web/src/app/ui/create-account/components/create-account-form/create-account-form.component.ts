import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
  selector: 'app-create-account-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './create-account-form.component.html',
  styleUrl: './create-account-form.component.css'
})
export class CreateAccountFormComponent {
  @Input() isSubmitting = false;
  @Output() createAccount = new EventEmitter<any>();

  createAccountForm: FormGroup;

  constructor(private fb: FormBuilder) {
    this.createAccountForm = this.fb.group({
      username: ['', [Validators.required, Validators.minLength(3)]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: ['', [Validators.required]]
    }, { validators: this.passwordMatchValidator });
  }

  passwordMatchValidator(form: FormGroup) {
    const password = form.get('password');
    const confirmPassword = form.get('confirmPassword');
    
    if (password && confirmPassword && password.value !== confirmPassword.value) {
      confirmPassword.setErrors({ passwordMismatch: true });
      return { passwordMismatch: true };
    }
    return null;
  }

  onSubmit(): void {
    if (this.createAccountForm.valid) {
      this.createAccount.emit(this.createAccountForm.value);
    } else {
      Object.keys(this.createAccountForm.controls).forEach(key => {
        this.createAccountForm.get(key)?.markAsTouched();
      });
    }
  }

  get username() {
    return this.createAccountForm.get('username');
  }

  get email() {
    return this.createAccountForm.get('email');
  }

  get password() {
    return this.createAccountForm.get('password');
  }

  get confirmPassword() {
    return this.createAccountForm.get('confirmPassword');
  }
}
