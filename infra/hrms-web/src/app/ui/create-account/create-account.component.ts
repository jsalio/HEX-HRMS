import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { CreateAccountFormComponent } from './components/create-account-form/create-account-form.component';

@Component({
  selector: 'app-create-account',
  standalone: true,
  imports: [CommonModule, CreateAccountFormComponent],
  template: `
    <app-create-account-form
      [isSubmitting]="isSubmitting"
      (createAccount)="onCreateAccount($event)">
    </app-create-account-form>
  `
})
export class CreateAccountComponent {
  isSubmitting = false;

  constructor(private router: Router) {}

  onCreateAccount(data: any): void {
    this.isSubmitting = true;
    console.log('Create account attempt:', data);
    
    // TODO: Usar el caso de uso una vez creado
    setTimeout(() => {
      this.isSubmitting = false;
      console.log('Account created successfully!');
      this.router.navigate(['/login']);
    }, 1500);
  }
}
