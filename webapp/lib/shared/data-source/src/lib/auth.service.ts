
import { HttpClient } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { Token, TokenApi, UserAuthForm } from "@arch-shared/types";
import { map, Observable } from "rxjs";
import { AuthService as ExternAuthService } from "@arch-shared/auth";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  readonly http = inject(HttpClient);
  readonly externAuthService = inject(ExternAuthService);

  readonly API_URL = '/api/auth';


  public login(form: UserAuthForm): Observable<Token> {
    return this.http.post<TokenApi>(`${this.API_URL}/login/internal`, form).pipe(
      map(token => this.populateTokenFromBackend(token)),
      this.externAuthService.storeToken()
    );
  }

  public refresh(): Observable<Token> {
    return this.http.get<TokenApi>(`${this.API_URL}/refresh`).pipe(
      map(token => this.populateTokenFromBackend(token)),
      this.externAuthService.storeToken()
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
