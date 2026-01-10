import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { LoginUserUseCase } from '../../core/usecases/login';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent implements OnInit {
  loginForm!: FormGroup;
  isSubmitting = false;

  constructor(private fb: FormBuilder, 
    private loginUseCase: LoginUserUseCase
  ) {}

  ngOnInit(): void {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      rememberMe: [false]
    });
  }

  onSubmit(): void {
    if (this.loginForm.valid) {
      this.isSubmitting = true;
      const { email, password, rememberMe } = this.loginForm.value;
      debugger
      // TODO: Implement actual authentication logic
      //console.log('Login attempt:', { email, password, rememberMe });
      this.loginUseCase.Execute({username:email, password}).then(() => {
        this.isSubmitting = false;
        console.log('Login successful!');
      }).catch(() => {
        this.isSubmitting = false;
        console.log('Login failed!');
      Object.keys(this.loginForm.controls).forEach(key => {
         this.loginForm.get(key)?.markAsTouched();
         });
      })
      // Simulate API call
    //   setTimeout(() => {
    //     this.isSubmitting = false;
    //     console.log('Login successful!');
    //   }, 1500);
    // } else {
    //   // Mark all fields as touched to show validation errors
    //   Object.keys(this.loginForm.controls).forEach(key => {
    //     this.loginForm.get(key)?.markAsTouched();
    //   });
    }
  }

  get email() {
    return this.loginForm.get('email');
  }

  get password() {
    return this.loginForm.get('password');
  }
}
