import { inject, Injectable } from "@angular/core";
import { AuthService } from "./auth.service";
import { map, switchMap } from "rxjs";
import { UserService } from "@arch-shared/data-source";
import { awakeColdObservable } from "@arch-shared/util";
import { toSignal } from "@angular/core/rxjs-interop";

@Injectable({
  providedIn: 'root'
})
export class CurrentUserService {

  readonly authService = inject(AuthService);
  readonly userService = inject(UserService);

  readonly currentUser$ = this.authService.isAuthenticated.pipe(
    switchMap(
      awakeColdObservable(this.userService.me.bind(this.userService))
    ),
    map(user => user || null) // Ensure we return null if user is not authenticated
  )

  readonly currentUser = toSignal(this.currentUser$, { initialValue: null });

}
