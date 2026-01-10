import { Component, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginFormComponent } from './components/login-form/login-form.component';
import { LoginUser } from '../../core/domain/models';
import { LoginUserUseCase } from '../../core/usecases/login';

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

  constructor(private loginUseCase: LoginUserUseCase) {}

  onLogin(user: LoginUser): void {
    this.isSubmitting.set(true);
    this.loginUseCase.Execute(user)
      .then(() => {
        this.isSubmitting.set(false);
        console.log('Login successful!');
        // Redirección aquí
      })
      .catch(() => {
        this.isSubmitting.set(false);
        console.log('Login failed!');
      });
  }
}
