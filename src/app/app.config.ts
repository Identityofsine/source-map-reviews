import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter, withComponentInputBinding } from '@angular/router';
import { providePrimeNG } from 'primeng/config';
import { routes } from './app.routes';
import { HTTP_INTERCEPTORS, provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { CookieInterceptor } from '../../lib/shared/data-source/src/lib/interceptors/CookieInterceptor';
import { ErrorInterceptor } from '../../lib/shared/data-source/src/lib/interceptors/ErrorInterceptor';
import { AuthInterceptor } from 'lib/shared/data-source/src/lib/interceptors/AuthInterceptor';
import Aura from '@primeuix/themes/aura';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes, withComponentInputBinding()),
    providePrimeNG({
      license: 'eyJpZCI6IjMwYjk4YTYzLTJkNTktNDA2ZS1iYTAzLWRlZmJhYjUzNmE1OSIsInByb2R1Y3QiOiJwcmltZXVpIiwidGllciI6ImNvbW11bml0eSIsInR5cGUiOiJkZXYiLCJpYXQiOjE3ODM3Mjg4OTgsImV4cCI6MTgxNTI2NDg5OH0._8_djlCuOb5KrcuHluv0-dwHOkcIF3F59nrLu4ssi7ujhnEEjB9R60M8HpGllVryr8YA5l8I6XwzGNgUc5NABg',
      theme: {
        preset: Aura,
        // Default options,
        options: {
          prefix: 'p',
          darkModeSelector: 'system',
          cssLayer: false,
          cssVariables: true
        }
      }
    }),
    provideHttpClient(withInterceptorsFromDi()),
    {
      provide: HTTP_INTERCEPTORS,
      useClass: CookieInterceptor,
      multi: true
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: ErrorInterceptor,
      multi: true
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    }
  ]
};
