import { Component, inject, signal, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { CreateUserUseCase } from '../../../../../core/usecases/create';
import { GetUserByFieldUseCase } from '../../../../../core/usecases/get-by-field';
import { ModifyUserUseCase } from '../../../../../core/usecases/modify';
import { CreateUser, ModifyUser, UserData } from '../../../../../core/domain/models';
import { UserFormComponent } from '../../components/user-form/user-form.component';

@Component({
  selector: 'app-user-create',
  standalone: true,
  imports: [CommonModule, UserFormComponent],
  templateUrl: './user-create.component.html'
})
export class UserCreateComponent implements OnInit {
  private createUserUseCase = inject(CreateUserUseCase);
  private getUserByFieldUseCase = inject(GetUserByFieldUseCase);
  private modifyUserUseCase = inject(ModifyUserUseCase);
  private router = inject(Router);
  private route = inject(ActivatedRoute);

  saving = signal(false);
  error = signal<string | null>(null);
  userToEdit = signal<UserData | null>(null);
  isEditing = signal(false);
  editId: string | null = null;

  ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.isEditing.set(true);
      this.editId = id;
      this.loadUser(id);
    }
  }

  loadUser(id: string) {
    this.getUserByFieldUseCase.Execute({ key: 'id', value: id })
      .then(user => {
        this.userToEdit.set(user);
      })
      .catch(err => {
        console.error('Error loading user:', err);
        this.error.set('Failed to load user data.');
      });
  }

  onSave(formData: any) {
    this.saving.set(true);
    this.error.set(null);

    if (this.isEditing() && this.editId) {
      const modifyUser: ModifyUser = {
        ...formData,
        id: this.editId
        // If password is empty in form, backend might overwrite? 
        // Logic depends on form handling. Assuming form sends values.
      };
      
      this.modifyUserUseCase.Execute(modifyUser)
        .then(() => {
          this.saving.set(false);
          this.router.navigate(['/settings']);
        })
        .catch((err) => {
          this.saving.set(false);
          console.error('Error updating user:', err);
          this.error.set('Failed to update user. Please try again.');
        });
    } else {
      const createUser: CreateUser = formData;
      this.createUserUseCase.Execute(createUser)
        .then(() => {
          this.saving.set(false);
          this.router.navigate(['/settings']);
        })
        .catch((err) => {
          this.saving.set(false);
          console.error('Error creating user:', err);
          this.error.set('Failed to create user. Please try again.');
        });
    }
  }

  onCancel() {
    this.router.navigate(['/settings']);
  }
}
