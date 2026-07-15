import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavComponent } from './nav/nav.component';
import { CurrentUserService } from 'lib/shared/auth/src/lib/current-user.service';
import { LookupsService } from '@arch-shared/lookups';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, NavComponent],
  providers: [LookupsService],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {

  readonly currentUserService = inject(CurrentUserService);
  readonly lookupsService = inject(LookupsService);

  title = 'webapp';

  constructor() {
    //load current user on app init
    this.currentUserService.currentUser();
  }
}
