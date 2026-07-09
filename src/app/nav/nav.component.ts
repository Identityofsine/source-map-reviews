import { ChangeDetectionStrategy, Component } from '@angular/core';
import { IconComponent } from '@arch-shared/arch-ui';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.scss'],
  imports: [
    IconComponent,
  ],
})
export class NavComponent {

}
