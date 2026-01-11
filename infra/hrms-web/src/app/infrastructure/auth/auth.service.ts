import { Injectable, signal, computed } from '@angular/core';
import { UserData } from '../../core/domain/models';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly AUTH_KEY = 'hrms_auth';
  
  // Signal to hold auth state
  private _currentUser = signal<UserData | null>(this.getStoredAuth());

  // Exposed read-only signal
  currentUser = computed(() => this._currentUser());
  isAuthenticated = computed(() => !!this._currentUser()?.token);

  constructor() {}

  setAuth(user: UserData): void {
    localStorage.setItem(this.AUTH_KEY, JSON.stringify(user));
    this._currentUser.set(user);
  }

  logout(): void {
    localStorage.removeItem(this.AUTH_KEY);
    this._currentUser.set(null);
  }

  private getStoredAuth(): UserData | null {
    const stored = localStorage.getItem(this.AUTH_KEY);
    if (!stored) return null;
    try {
      return JSON.parse(stored) as UserData;
    } catch {
      return null;
    }
  }

  getToken(): string | null {
    return this._currentUser()?.token || null;
  }
}
