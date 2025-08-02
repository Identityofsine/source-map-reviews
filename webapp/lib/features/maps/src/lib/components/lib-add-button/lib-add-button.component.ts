import { ChangeDetectionStrategy, Component, output } from '@angular/core';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-add-button',
  templateUrl: './lib-add-button.component.html',
  styleUrls: ['./lib-add-button.component.scss'],
  imports: [],
})
export class AddButtonComponent {

  readonly click = output<void>();

}
