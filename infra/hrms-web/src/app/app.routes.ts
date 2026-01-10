import { Routes } from '@angular/router';
import { authGuard } from './infrastructure/auth/auth.guard';

export const routes: Routes = [
  // Authentication routes (outside layout)
  {
    path: 'login',
    loadComponent: () => import('./ui/login/login.component').then(m => m.LoginComponent)
  },
  {
    path: 'create-account',
    loadComponent: () => import('./ui/create-account/create-account.component').then(m => m.CreateAccountComponent)
  },
  
  // Main application routes (inside layout)
  {
    path: '',
    canActivate: [authGuard],
    loadComponent: () => import('./ui/shared/layout/layout.component').then(m => m.LayoutComponent),
    children: [
      {
        path: 'dashboard',
        loadComponent: () => import('./ui/site/home/home.component').then(m => m.HomeComponent)
      },
      {
        path: 'employees',
        loadComponent: () => import('./ui/site/home/home.component').then(m => m.HomeComponent) // TODO: Create EmployeesComponent
      },
      {
        path: 'departments',
        loadComponent: () => import('./ui/site/home/home.component').then(m => m.HomeComponent) // TODO: Create DepartmentsComponent
      },
      {
        path: 'positions',
        loadComponent: () => import('./ui/site/home/home.component').then(m => m.HomeComponent) // TODO: Create PositionsComponent
      },
      {
        path: 'attendance',
        loadComponent: () => import('./ui/site/home/home.component').then(m => m.HomeComponent) // TODO: Create AttendanceComponent
      },
      {
        path: 'payroll',
        loadComponent: () => import('./ui/site/home/home.component').then(m => m.HomeComponent) // TODO: Create PayrollComponent
      },
      {
        path: 'leave-requests',
        loadComponent: () => import('./ui/site/home/home.component').then(m => m.HomeComponent) // TODO: Create LeaveRequestsComponent
      },
      {
        path: 'settings',
        loadComponent: () => import('./ui/site/settings/settings.component').then(m => m.SettingsComponent)
      },
      {
        path: '',
        redirectTo: 'login',
        pathMatch: 'full'
      }
    ]
  }
];
