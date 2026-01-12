import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListUserUseCase } from '../../../core/usecases/list';
import { ListRoleUseCase } from '../../../core/usecases/role';
import { UserData, Role } from '../../../core/domain/models';
import { UserListComponent } from './components/user-list/user-list.component';
import { RoleListComponent } from './components/role-list/role-list.component';
import { RoleFormComponent } from './components/role-form/role-form.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [CommonModule, UserListComponent, RoleListComponent, RoleFormComponent],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.css'
})
export class SettingsComponent implements OnInit {
  private listUser = inject(ListUserUseCase);
  private listRole = inject(ListRoleUseCase);
  private router = inject(Router);

  activeTab = signal<'users' | 'roles'>('users');
  users = signal<UserData[]>([]);
  roles = signal<Role[]>([]);
  currentPage = signal<number>(1);
  totalPages = signal<number>(1);
  pageSize = signal<number>(5);

  // Modal State
  showRoleForm = signal<boolean>(false);
  selectedRole = signal<Role | null>(null);

  ngOnInit(): void {
    this.fetchUsers();
    this.fetchRoles();
  }

  fetchUsers(page: number = 1): void {
    this.listUser.Execute({ 
      filters: [], 
      pagination: { page: page, limit: this.pageSize() } 
    }).then((data) => {
      this.users.set(data.rows);
      this.totalPages.set(data.total_pages);
      this.currentPage.set(page);
    }).catch(err => {
      console.error('Error fetching users:', err);
    });
  }

  fetchRoles(): void {
    this.listRole.Execute({
      filters: [],
      pagination: { page: 1, limit: 100 }
    }).then((data) => {
      this.roles.set(data.rows);
    }).catch(err => {
        console.error('Error fetching roles:', err);
    });
  }

  changePage(page: number): void {
    if (page >= 1 && page <= this.totalPages()) {
        this.fetchUsers(page);
    }
  }

  setTab(tab: 'users' | 'roles') {
    this.activeTab.set(tab);
  }

  deleteUser(id: string) {
    console.log('Delete user', id);
    // TODO: Implement real delete
    this.users.update(users => users.filter(u => u.id !== id));
  }
  editUser(id: string) {
    this.router.navigate(['settings/users/edit', id]);
  }

  toggleStatus(id: string) {
    console.log('Toggle status', id);
  }

  // --- Role Actions ---

  createRole() {
    this.selectedRole.set(null);
    this.showRoleForm.set(true);
  }

  editRole(role: Role) {
    this.selectedRole.set(role);
    this.showRoleForm.set(true);
  }

  closeRoleForm() {
    this.showRoleForm.set(false);
    this.selectedRole.set(null);
  }

  onRoleSaved() {
    this.fetchRoles();
    this.closeRoleForm();
  }
}
