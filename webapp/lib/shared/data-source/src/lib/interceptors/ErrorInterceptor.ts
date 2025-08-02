import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { inject, Injectable, Injector } from '@angular/core';
import { catchError, Observable, switchMap, throwError } from 'rxjs';
import { AuthService } from '../auth.service';
import { AuthService as ExternAuthService } from '@arch-shared/auth';
import { Router } from '@angular/router';

@Injectable()
export class ErrorInterceptor implements HttpInterceptor {

  //DI
  readonly injector = inject(Injector);
  readonly router = inject(Router);

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {

    return next.handle(req).pipe(
      catchError((error: HttpErrorResponse) => {

        const authService = this.injector.get(AuthService); // Ensure AuthService is injected to handle token refresh
        const externAuthService = this.injector.get(ExternAuthService); // Ensure ExternAuthService is injected to handle token management

        // Handle specific status codes
        switch (error.status) {
          case 419:
            return authService.refresh().pipe(
              switchMap(() => next.handle(req)),
              catchError((refreshError) => {
                // Handle refresh token failure, e.g., redirect to login
                externAuthService.wipeToken();
                return throwError(() => refreshError);
              }));
          case 422:
            externAuthService.wipeToken();
            return throwError(() => error);
          case 401:
          case 403:
          default:
            return throwError(() => error);
        }
      }));
  }

}

