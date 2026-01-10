import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterModule, RouterOutlet } from '@angular/router';
import { AuthService } from '../../../infrastructure/auth/auth.service';

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
  private authService = inject(AuthService);
  private router = inject(Router);

  isSidebarCollapsed = false;
  currentUser = this.authService.currentUser;

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
    this.authService.logout();
    this.router.navigate(['/login']);
  }
}
