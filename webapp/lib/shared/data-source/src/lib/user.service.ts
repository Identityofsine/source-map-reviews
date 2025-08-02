import { HttpClient } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { Token, TokenApi, UserAuthForm } from "@arch-shared/types";
import { map, Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class UserService {

  readonly http = inject(HttpClient);

  readonly API_URL = '/api/user';
  readonly AUTH_URL = '/api/auth';


  public login(form: UserAuthForm): Observable<Token> {
    return this.http.post<TokenApi>(`${this.AUTH_URL}/internal`, form).pipe(
      map(token => this.populateTokenFromBackend(token))
    );

  }

  private populateTokenFromBackend(token: TokenApi): Token {
    return {
      ...token,
      expiresAt: token.expiresAt ? new Date(token.expiresAt) : undefined,
      createdAt: token.createdAt ? new Date(token.createdAt) : undefined,
      updatedAt: token.updatedAt ? new Date(token.updatedAt) : undefined
    };
  }

}
