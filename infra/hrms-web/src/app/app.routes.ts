import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: 'login',
    loadComponent: () => import('./ui/login/login.component').then(m => m.LoginComponent)
  },
  {
    path: 'create-account',
    loadComponent: () => import('./ui/create-account/create-account.component').then(m => m.CreateAccountComponent)
  },
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  }
];
