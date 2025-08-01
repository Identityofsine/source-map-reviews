import { ChangeDetectionStrategy, Component, computed, inject, input } from '@angular/core';
import { MapTag } from '@arch-shared/types';
import { Router } from '@angular/router';
import { MapTagComponent } from '@arch-feature/maps';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-thumbnail',
  templateUrl: './lib-map-thumbnail.component.html',
  styleUrls: ['./lib-map-thumbnail.component.scss'],
  imports: [
    MapTagComponent
  ],
})
export class MapThumbnailComponent {

  // DI
  readonly router = inject(Router);

  readonly mapName = input<string>();
  readonly mapTags = input<MapTag[]>();

  readonly mapImage = computed(() => {
    return `/api/images/${this.mapName() ?? 'map_placeholder'}.jpg`;
  })

  onImgError(event: Event) {
    (event.target as HTMLImageElement).src = '/maps/map_placeholder.jpg';
  }

  onMapClick() {
    this.router.navigate(['/map', this.mapName()]);
  }

}
