import { inject, Injectable } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { UserService } from '@arch-shared/data-source';
import { Token } from '@arch-shared/types';
import { BehaviorSubject, map, MonoTypeOperatorFunction, Observable, of, skip, Subscription, catchError, switchMap } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  readonly userService = inject(UserService);
  readonly storage$ = new BehaviorSubject<Token | null>(this.getToken());
  readonly isValidated$ = new BehaviorSubject<boolean>(false);

  readonly localStorageSubscription: Subscription;
  readonly isAuthenticated = this.storage$.pipe(
    switchMap(token => {
      if (!token || !token.accessToken || new Date(token.expiresAt) < new Date()) {
        return of(false);
      }
      // Just check if we have a valid token locally to avoid circular dependency
      return this.userService.validateInitial(token.accessToken).pipe(
        map(user => !!user && !!user.id),
        catchError(() => {
          console.warn('User validation failed, token might be invalid');
          return of(false)
        })
      )
    }),
  );

  readonly isAuthenticatedSignal = toSignal(this.isAuthenticated, { initialValue: false });

  constructor() {
    this.storage$.next(this.getToken());
    this.localStorageSubscription = (
      this.storage$.subscribe(token => {
        if (token) {
          localStorage.setItem('token', JSON.stringify(token));
        } else {
          localStorage.removeItem('token');
        }
      })
    );

  }

  public storeToken(): MonoTypeOperatorFunction<Token> {
    return (source: Observable<Token>) => new Observable<Token>(subscriber => {
      source.subscribe({
        next: (token) => {
          if (!token || !token?.accessToken) {
            subscriber.error(new Error('Token is null or undefined'));
            return;
          }
          this.storage$.next(token);
          subscriber.next(token);
        },
        error: (err) => subscriber.error(err),
        complete: () => subscriber.complete()
      });
    });
  }

  public wipeToken(): void {
    try {
      this.storage$.next(null);
    } catch (error) {
      console.error('Error removing token from localStorage:', error);
    }
  }

  public getToken(): Token | null {
    try {
      const token = localStorage.getItem('token');
      return token ? JSON.parse(token) : null;
    } catch (error) {
      return null;
    }
  }

  // Helper method if you need to validate the user with the server
  public validateUserWithServer(): Observable<boolean> {
    const userService = inject(UserService);
    return userService.me().pipe(
      map(user => !!user && !!user.id)
    );
  }

}
