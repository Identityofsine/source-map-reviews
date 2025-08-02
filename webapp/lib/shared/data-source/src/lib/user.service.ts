import { HttpClient } from "@angular/common/http";
import { inject, Injectable, Injector } from "@angular/core";
import { User } from "@arch-shared/types";
import { map, Observable } from "rxjs";

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
