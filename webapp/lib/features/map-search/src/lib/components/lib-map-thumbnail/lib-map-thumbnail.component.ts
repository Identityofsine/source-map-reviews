import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapTag } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-thumbnail',
  templateUrl: './lib-map-thumbnail.component.html',
  styleUrls: ['./lib-map-thumbnail.component.scss'],
  imports: [],
})
export class MapThumbnailComponent {

  readonly mapName = input<string>();
  readonly mapTags = input<MapTag[]>();

}
