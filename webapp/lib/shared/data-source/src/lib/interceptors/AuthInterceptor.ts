import {
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
} from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { AuthService } from '@arch-shared/auth';
import { Observable } from 'rxjs';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  readonly AuthService = inject(AuthService);

  intercept(
    req: HttpRequest<any>,
    next: HttpHandler,
  ): Observable<HttpEvent<any>> {

    const token = this.AuthService.getToken();
    if (!token || !token.accessToken) {
      return next.handle(req);
    }
    const authHeader = `Bearer ${token.accessToken}`;


    const clonedRequest = req.clone({
      headers: req.headers.set('Authorization', authHeader),
    });

    return next.handle(clonedRequest);
  }

}
