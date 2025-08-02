import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, Observable, switchMap, throwError } from 'rxjs';
import { AuthService } from '../auth.service';
import { AuthService as ExternAuthService } from '@arch-shared/auth';
import { Router } from '@angular/router';

@Injectable()
export class ErrorInterceptor implements HttpInterceptor {

  //DI
  readonly authService = inject(AuthService);
  readonly externAuthService = inject(ExternAuthService);
  readonly router = inject(Router);

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(
      catchError((error: HttpErrorResponse) => {
        // Handle specific status codes
        switch (error.status) {
          case 419:
            return this.authService.refresh().pipe(
              switchMap(() => next.handle(req)),
              catchError((refreshError) => {
                // Handle refresh token failure, e.g., redirect to login
                this.externAuthService.wipeToken();
                return throwError(() => refreshError);
              }));
          case 422:
            this.externAuthService.wipeToken();
            return throwError(() => error);
          case 401:
          case 403:
          default:
            return throwError(() => error);
        }
      }));
  }

}

