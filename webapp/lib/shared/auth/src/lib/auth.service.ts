import { Injectable } from '@angular/core';
import { Token } from '@arch-shared/types';
import { MonoTypeOperatorFunction, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  public storeToken(): MonoTypeOperatorFunction<Token> {
    return (source: Observable<Token>) => new Observable<Token>(subscriber => {
      source.subscribe({
        next: (token) => {
          if (!token || !token?.accessToken) {
            subscriber.error(new Error('Token is null or undefined'));
            return;
          }
          localStorage.setItem('token', JSON.stringify(token));
          subscriber.next(token);
        },
        error: (err) => subscriber.error(err),
        complete: () => subscriber.complete()
      });
    });
  }

  public wipeToken(): void {
    try {
      localStorage.removeItem('token');
    } catch (error) {
      console.error('Error removing token from localStorage:', error);
    }
  }

  public isUserAuthenticated(): boolean {
    const token = this.getToken();
    return !!token && !!token.accessToken && new Date(token.expiresAt) > new Date();
  }

  public getToken(): Token | null {
    try {
      const token = localStorage.getItem('token');
      return token ? JSON.parse(token) : null;
    } catch (error) {
      return null;
    }
  }


}
