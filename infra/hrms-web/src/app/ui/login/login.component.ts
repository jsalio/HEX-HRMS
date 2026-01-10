import { Component, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginFormComponent } from './components/login-form/login-form.component';
import { LoginUser } from '../../core/domain/models';
import { LoginUserUseCase } from '../../core/usecases/login';
import { AuthService } from '../../infrastructure/auth/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, LoginFormComponent],
  template: `
    <app-login-form 
      [isSubmitting]="isSubmitting()" 
      (login)="onLogin($event)">
    </app-login-form>
  `
})
export class LoginComponent {
  isSubmitting = signal<boolean>(false);

  constructor(
    private loginUseCase: LoginUserUseCase,
    private authService: AuthService,
    private router: Router
  ) {}

  onLogin(user: LoginUser): void {
    this.isSubmitting.set(true);
    this.loginUseCase.Execute(user)
      .then((userData) => {
        debugger
        this.isSubmitting.set(false);
        this.authService.setAuth(userData);
        console.log('Login successful!');
        this.router.navigate(['/dashboard']);
      })
      .catch(() => {
        this.isSubmitting.set(false);
        console.log('Login failed!');
      });
  }
}
