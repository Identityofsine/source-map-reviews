import { ChangeDetectionStrategy, Component, computed, inject, input } from '@angular/core';
import { MapTag } from '@arch-shared/types';
import { Router } from '@angular/router';
import { MapTagsComponent } from '@arch-feature/maps';
import { AuthService } from '@arch-shared/auth';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-thumbnail',
  templateUrl: './lib-map-thumbnail.component.html',
  styleUrls: ['./lib-map-thumbnail.component.scss'],
  imports: [
    MapTagsComponent
  ],
})
export class MapThumbnailComponent {

  // DI
  readonly router = inject(Router);
  readonly authService = inject(AuthService);

  readonly mapName = input<string>();
  readonly mapTags = input<MapTag[]>();

  readonly mapImage = computed(() => {
    return `/api/images/${this.mapName() ?? 'map_placeholder'}.jpg`;
  })

  readonly isAuthenticated = this.authService.isAuthenticatedSignal;

  onImgError(event: Event) {
    (event.target as HTMLImageElement).src = '/maps/map_placeholder.jpg';
  }

  onMapClick() {
    this.router.navigate(['/map', this.mapName()]);
  }

}
