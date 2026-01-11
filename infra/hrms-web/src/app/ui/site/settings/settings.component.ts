import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListUserUseCase } from '../../../core/usecases/list';
import { UserData } from '../../../core/domain/models';
import { UserListComponent } from './components/user-list/user-list.component';
import { Router } from '@angular/router';

interface RoleMock {
  id: string;
  name: string;
  permissions: string[];
}

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [CommonModule, UserListComponent],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.css'
})
export class SettingsComponent implements OnInit {
  private list = inject(ListUserUseCase);
  private router = inject(Router);

  activeTab = signal<'users' | 'roles'>('users');
  users = signal<UserData[]>([]);
  currentPage = signal<number>(1);
  totalPages = signal<number>(1);
  pageSize = signal<number>(5);
  roles = signal<RoleMock[]>([
    { id: '1', name: 'Administrator', permissions: ['all_access', 'manage_users', 'manage_roles'] },
    { id: '2', name: 'Manager', permissions: ['view_reports', 'manage_employees'] },
    { id: '3', name: 'Developer', permissions: ['view_dashboard', 'manage_code'] },
  ]);

  ngOnInit(): void {
    this.fetchUsers();
  }

  fetchUsers(page: number = 1): void {
    this.list.Execute({ 
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
    console.log('Edit user', id);
    // TODO: Implement real edit
    // this.users.update(users => users.filter(u => u.id !== id));
    this.router.navigate(['settings/users/edit', id]);
  }

  toggleStatus(id: string) {
    console.log('Toggle status', id);
    // TODO: Implement real status toggle
  }
}
