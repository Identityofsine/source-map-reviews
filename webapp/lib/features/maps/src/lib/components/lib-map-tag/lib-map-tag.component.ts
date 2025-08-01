import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MapTag } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-tag',
  templateUrl: './lib-map-tag.component.html',
  styleUrls: ['./lib-map-tag.component.scss'],
  imports: [],
})
export class MapTagComponent {

  readonly tag = input<MapTag>();
  readonly click = output<MapTag>();

}
