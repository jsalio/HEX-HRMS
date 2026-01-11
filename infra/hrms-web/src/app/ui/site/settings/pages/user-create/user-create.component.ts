import { Component, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { CreateUserUseCase } from '../../../../../core/usecases/create';
import { CreateUser } from '../../../../../core/domain/models';
import { UserFormComponent } from '../../components/user-form/user-form.component';

@Component({
  selector: 'app-user-create',
  standalone: true,
  imports: [CommonModule, UserFormComponent],
  templateUrl: './user-create.component.html'
})
export class UserCreateComponent {
  private createUserUseCase = inject(CreateUserUseCase);
  private router = inject(Router);

  saving = signal(false);
  error = signal<string | null>(null);

  onSave(user: CreateUser) {
    this.saving.set(true);
    this.error.set(null);

    this.createUserUseCase.Execute(user)
      .then(() => {
        this.saving.set(false);
        this.router.navigate(['/settings']);
      })
      .catch((err) => {
        this.saving.set(false);
        console.error('Error creating user:', err);
        // Extract error message if available, otherwise generic
        this.error.set('Failed to create user. Please try again.');
      });
  }

  onCancel() {
    this.router.navigate(['/settings']);
  }
}
