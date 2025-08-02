import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { IconComponent } from '@arch-shared/arch-ui';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-add-button',
  templateUrl: './lib-add-button.component.html',
  styleUrls: ['./lib-add-button.component.scss'],
  imports: [
    IconComponent
  ],
})
export class AddButtonComponent {


  readonly shouldShow = input<boolean>(false);
  readonly click = output<void>();

}
