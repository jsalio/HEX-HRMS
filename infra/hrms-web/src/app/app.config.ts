import { ApplicationConfig, InjectionToken, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { UserRepository } from './core/domain/ports/user.repo';
import { provideHttpClient } from '@angular/common/http';

export const USER_REPOSITORY =
  new InjectionToken<UserRepository>('USER_REPOSITORY');

export const appConfig: ApplicationConfig = {
  providers: [provideZoneChangeDetection({ eventCoalescing: true }), provideRouter(routes), provideHttpClient()]
};
