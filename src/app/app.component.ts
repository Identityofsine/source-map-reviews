import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavComponent } from './nav/nav.component';
import { CurrentUserService } from 'lib/shared/auth/src/lib/current-user.service';
import { LookupCacheService } from '@arch-shared/data-source';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, NavComponent],
  providers: [LookupCacheService],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {

  readonly currentUserService = inject(CurrentUserService);
  readonly lookupsService = inject(LookupCacheService);

  readonly lookupsLoading = this.lookupsService.loading;

  title = 'webapp';

  constructor() {
    //load current user on app init
    this.currentUserService.currentUser();
  }
}
