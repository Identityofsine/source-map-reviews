import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapTag } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-header',
  templateUrl: './lib-map-header.component.html',
  styleUrls: ['./lib-map-header.component.scss'],
  imports: [],
})
export class MapHeaderComponent {

  readonly mapName = input<string>();
  readonly mapTags = input<MapTag[]>();

}
