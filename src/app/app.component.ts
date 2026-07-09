import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavComponent } from './nav/nav.component';
import { CurrentUserService } from 'lib/shared/auth/src/lib/current-user.service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, NavComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {

  readonly currentUserService = inject(CurrentUserService);

  title = 'webapp';

  constructor() {
    //load current user on app init
    this.currentUserService.currentUser();
  }
}
