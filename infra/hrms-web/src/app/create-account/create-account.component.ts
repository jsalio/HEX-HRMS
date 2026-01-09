import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-account',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './create-account.component.html',
  styleUrl: './create-account.component.css'
})
export class CreateAccountComponent implements OnInit {
  createAccountForm!: FormGroup;
  isSubmitting = false;

  constructor(
    private fb: FormBuilder,
    private router: Router
  ) {}

  ngOnInit(): void {
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
      this.isSubmitting = true;
      const { username, email, password } = this.createAccountForm.value;
      
      // TODO: Implement actual account creation logic
      console.log('Create account attempt:', { username, email, password });
      
      // Simulate API call
      setTimeout(() => {
        this.isSubmitting = false;
        console.log('Account created successfully!');
        // Navigate to login after successful creation
        // this.router.navigate(['/login']);
      }, 1500);
    } else {
      // Mark all fields as touched to show validation errors
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
