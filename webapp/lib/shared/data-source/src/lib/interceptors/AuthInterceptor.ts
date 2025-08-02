import {
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
} from '@angular/common/http';
import { inject, Injectable, Injector } from '@angular/core';
import { AuthService } from '@arch-shared/auth';
import { defer, Observable } from 'rxjs';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  readonly injector = inject(Injector);

  intercept(
    req: HttpRequest<any>,
    next: HttpHandler,
  ): Observable<HttpEvent<any>> {

    return defer(() => {
      // Check for bypass header to avoid circular dependency during initial validation
      if (req.headers.has('X-Skip-Auth-Interceptor')) {
        return next.handle(req);
      }

      const authService = this.injector.get(AuthService);
      const token = authService.getToken();

      if (!token || !token.accessToken) {
        return next.handle(req);
      }

      const authHeader = `Bearer ${token.accessToken}`;
      const clonedRequest = req.clone({
        headers: req.headers.set('Authorization', authHeader),
      });

      return next.handle(clonedRequest);
    });
  }

}
