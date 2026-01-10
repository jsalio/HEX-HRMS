import { Component, signal } from '@angular/core';
import { CommonModule } from '@angular/common';

interface UserMock {
  id: string;
  name: string;
  email: string;
  role: string;
  status: 'active' | 'inactive';
}

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
export class SettingsComponent {
  activeTab = signal<'users' | 'roles'>('users');

  users = signal<UserMock[]>([
    { id: '1', name: 'Admin User', email: 'admin@hex.com', role: 'Administrator', status: 'active' },
    { id: '2', name: 'Jorge Salio', email: 'jorge@hex.com', role: 'Developer', status: 'active' },
    { id: '3', name: 'Demo User', email: 'demo@hex.com', role: 'Manager', status: 'inactive' },
  ]);

  roles = signal<RoleMock[]>([
    { id: '1', name: 'Administrator', permissions: ['all_access', 'manage_users', 'manage_roles'] },
    { id: '2', name: 'Manager', permissions: ['view_reports', 'manage_employees'] },
    { id: '3', name: 'Developer', permissions: ['view_dashboard', 'manage_code'] },
  ]);

  setTab(tab: 'users' | 'roles') {
    this.activeTab.set(tab);
  }

  deleteUser(id: string) {
    console.log('Delete user', id);
    this.users.update(users => users.filter(u => u.id !== id));
  }

  toggleStatus(id: string) {
    console.log('Toggle status', id);
    this.users.update(users => users.map(u => 
      u.id === id ? { ...u, status: u.status === 'active' ? 'inactive' : 'active' } : u
    ));
  }
}
