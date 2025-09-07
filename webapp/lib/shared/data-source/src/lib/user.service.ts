import { HttpClient } from "@angular/common/http";
import { inject, Injectable, Injector } from "@angular/core";
import { User } from "@arch-shared/types";
import { catchError, map, Observable, of } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class UserService {

  readonly injector = inject(Injector);

  private get http(): HttpClient {
    return this.injector.get(HttpClient);
  }

  readonly API_URL = '/api/user';

  public me(): Observable<User> {
    return this.http.get<User>(`${this.API_URL}/me`)
  }

  public getUsername(userId: number): Observable<string> {
    return this.http.get<User>(`${this.API_URL}/${userId}`).pipe(
      catchError((err: any) => {
        // Catch all errors and return null, which will map to 'Unknown'
        return of(null);
      }),
      map(user => user?.details?.firstName || 'Unknown')
    );
  }


  // Special method for initial validation that bypasses auth interceptor
  public validateInitial(token: string): Observable<User> {
    return this.http.get<User>(`${this.API_URL}/me`, {
      headers: {
        'X-Skip-Auth-Interceptor': 'true',
        'Authorization': `Bearer ${token}`
      }
    });
  }

}
