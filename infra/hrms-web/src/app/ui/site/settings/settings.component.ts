import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListUserUseCase } from '../../../core/usecases/list';
import { ListRoleUseCase } from '../../../core/usecases/role';
import { ListDepartmentUseCase, DeleteDepartmentUseCase } from '../../../core/usecases/department';
import { UserData, Role, Department } from '../../../core/domain/models';
import { UserListComponent } from './components/user-list/user-list.component';
import { RoleListComponent } from './components/role-list/role-list.component';
import { RoleFormComponent } from './components/role-form/role-form.component';
import { DepartmentListComponent } from '../department/components/department-list/department-list.component';
import { DepartmentFormComponent } from '../department/components/department-form/department-form.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [CommonModule, UserListComponent, RoleListComponent, RoleFormComponent, DepartmentListComponent, DepartmentFormComponent],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.css'
})
export class SettingsComponent implements OnInit {
  private listUser = inject(ListUserUseCase);
  private listRole = inject(ListRoleUseCase);
  private listDepartment = inject(ListDepartmentUseCase);
  private deleteDepartment = inject(DeleteDepartmentUseCase);
  private router = inject(Router);

  activeTab = signal<'users' | 'roles' | 'departments'>('users');
  users = signal<UserData[]>([]);
  roles = signal<Role[]>([]);
  departments = signal<Department[]>([]);
  currentPage = signal<number>(1);
  totalPages = signal<number>(1);
  pageSize = signal<number>(5);

  // Role Modal State
  showRoleForm = signal<boolean>(false);
  selectedRole = signal<Role | null>(null);

  // Department Modal State
  showDepartmentForm = signal<boolean>(false);
  selectedDepartment = signal<Department | null>(null);

  ngOnInit(): void {
    this.fetchUsers();
    this.fetchRoles();
    this.fetchDepartments();
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

  fetchDepartments(): void {
    this.listDepartment.Execute({
      filters: [],
      pagination: { page: 1, limit: 100 }
    }).then((data) => {
      this.departments.set(data.rows);
    }).catch(err => {
        console.error('Error fetching departments:', err);
    });
  }

  changePage(page: number): void {
    if (page >= 1 && page <= this.totalPages()) {
        this.fetchUsers(page);
    }
  }

  setTab(tab: 'users' | 'roles' | 'departments') {
    this.activeTab.set(tab);
  }

  deleteUser(id: string) {
    console.log('Delete user', id);
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

  // --- Department Actions ---

  createDepartment() {
    this.selectedDepartment.set(null);
    this.showDepartmentForm.set(true);
  }

  editDepartment(department: Department) {
    this.selectedDepartment.set(department);
    this.showDepartmentForm.set(true);
  }

  deleteDepartmentAction(id: string) {
    this.deleteDepartment.Execute(id).then(() => {
      this.fetchDepartments();
    }).catch(err => {
      console.error('Error deleting department:', err);
    });
  }

  closeDepartmentForm() {
    this.showDepartmentForm.set(false);
    this.selectedDepartment.set(null);
  }

  onDepartmentSaved() {
    this.fetchDepartments();
    this.closeDepartmentForm();
  }
}
