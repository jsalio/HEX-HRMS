import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListUserUseCase } from '../../../core/usecases/list';
import { UserData } from '../../../core/domain/models';

interface RoleMock {
  id: string;
  name: string;
  permissions: string[];
}

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.css'
})
export class SettingsComponent implements OnInit {
  private list = inject(ListUserUseCase);

  activeTab = signal<'users' | 'roles'>('users');
  users = signal<UserData[]>([]);
  roles = signal<RoleMock[]>([
    { id: '1', name: 'Administrator', permissions: ['all_access', 'manage_users', 'manage_roles'] },
    { id: '2', name: 'Manager', permissions: ['view_reports', 'manage_employees'] },
    { id: '3', name: 'Developer', permissions: ['view_dashboard', 'manage_code'] },
  ]);

  ngOnInit(): void {
    this.fetchUsers();
  }

  fetchUsers(): void {
    this.list.Execute({ key: '', value: '' }).then((data) => {
      this.users.set(data);
    }).catch(err => {
      console.error('Error fetching users:', err);
    });
  }

  setTab(tab: 'users' | 'roles') {
    this.activeTab.set(tab);
  }

  getInitials(user: UserData): string {
    if (!user.name || !user.lastName) return '??';
    return (user.name[0] + user.lastName[0]).toUpperCase();
  }

  deleteUser(id: string) {
    console.log('Delete user', id);
    // TODO: Implement real delete
    this.users.update(users => users.filter(u => u.id !== id));
  }

  toggleStatus(id: string) {
    console.log('Toggle status', id);
    // TODO: Implement real status toggle
  }
}
