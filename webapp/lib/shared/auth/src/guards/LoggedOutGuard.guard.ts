import { inject, Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { map, of, take, tap } from 'rxjs';
import { AuthService } from '../lib/auth.service';

@Injectable({
  providedIn: 'root',
})
export class LoggedOutGuard implements CanActivate {
  private authService = inject(AuthService);
  private router = inject(Router);

  constructor() { }

  canActivate(_: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    return of(this.authService.isUserAuthenticated()).pipe(
      take(1), // Complete the observable after first emission
      map((isLoggedIn) => {
        if (isLoggedIn) {
          // Return a UrlTree for redirect instead of navigating directly
          return this.router.createUrlTree(['/app']);
        } else {
          return true;
        }
      }),
    );
  }
}
