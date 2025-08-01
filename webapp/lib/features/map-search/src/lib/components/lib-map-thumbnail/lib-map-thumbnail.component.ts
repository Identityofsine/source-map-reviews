import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';
import { MapTag } from '@arch-shared/types';
import { MapThumbnailTagComponent } from '../lib-map-thumbnail-tag/lib-map-thumbnail-tag.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-thumbnail',
  templateUrl: './lib-map-thumbnail.component.html',
  styleUrls: ['./lib-map-thumbnail.component.scss'],
  imports: [
    MapThumbnailTagComponent
  ],
})
export class MapThumbnailComponent {

  readonly mapName = input<string>();
  readonly mapTags = input<MapTag[]>();

  readonly mapImage = computed(() => {
    return `/api/images/${this.mapName() ?? 'map_placeholder'}.jpg`;
  })

  onImgError(event: Event) {
    (event.target as HTMLImageElement).src = '/maps/map_placeholder.jpg';
  }

}
