import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, RouterOutlet } from '@angular/router';

interface MenuItem {
  icon: string;
  label: string;
  route: string;
  badge?: number;
}

@Component({
  selector: 'app-layout',
  standalone: true,
  imports: [CommonModule, RouterModule, RouterOutlet],
  templateUrl: './layout.component.html',
  styleUrl: './layout.component.css'
})
export class LayoutComponent {
  isSidebarCollapsed = false;
  currentUser = {
    name: 'John Doe',
    role: 'Administrator',
    avatar: 'JD'
  };

  menuItems: MenuItem[] = [
    { icon: '📊', label: 'Dashboard', route: '/dashboard' },
    { icon: '👥', label: 'Employees', route: '/employees' },
    { icon: '🏢', label: 'Departments', route: '/departments' },
    { icon: '💼', label: 'Positions', route: '/positions' },
    { icon: '📅', label: 'Attendance', route: '/attendance' },
    { icon: '💰', label: 'Payroll', route: '/payroll' },
    { icon: '📝', label: 'Leave Requests', route: '/leave-requests', badge: 3 },
    { icon: '⚙️', label: 'Settings', route: '/settings' },
  ];

  toggleSidebar(): void {
    this.isSidebarCollapsed = !this.isSidebarCollapsed;
  }

  logout(): void {
    // TODO: Implement logout logic
    console.log('Logout clicked');
  }
}
