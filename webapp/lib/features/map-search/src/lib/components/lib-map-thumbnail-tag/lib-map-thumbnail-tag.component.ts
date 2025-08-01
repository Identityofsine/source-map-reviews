import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MapTag } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-thumbnail-tag',
  templateUrl: './lib-map-thumbnail-tag.component.html',
  styleUrls: ['./lib-map-thumbnail-tag.component.scss'],
  imports: [],
})
export class MapThumbnailTagComponent {

  readonly tag = input<MapTag>();
  readonly click = output<MapTag>();

}
