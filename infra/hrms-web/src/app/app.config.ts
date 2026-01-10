import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideHttpClient } from '@angular/common/http';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { UserApiRepository } from './infrastructure/userApiRepository.service';
import { USER_REPOSITORY } from './core/domain/ports/user.repo';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }), 
    provideRouter(routes), 
    provideHttpClient(),
    { provide: USER_REPOSITORY, useClass: UserApiRepository }
  ]
};
