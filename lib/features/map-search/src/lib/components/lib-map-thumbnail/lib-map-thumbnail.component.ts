import { ChangeDetectionStrategy, Component, computed, inject, input } from '@angular/core';
import { Map, MapApi, MapTag } from '@arch-shared/types';
import { Router } from '@angular/router';
import { MapTagsComponent } from '@arch-feature/maps';
import { AuthService } from '@arch-shared/auth';
import { MapSearchFormService } from '../../map-search-form.service';
import { getShowcaseMapImage } from '@arch-shared/util';

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
  readonly formService = inject(MapSearchFormService);
  readonly router = inject(Router);
  readonly authService = inject(AuthService);

  readonly map = input<MapApi>();

  readonly mapName = computed(() => {
    return this.map()?.name || '';
  });

  readonly mapTags = computed(() => {
    return this.map().categories || [];
  })

  readonly mapImage = computed(() => {
    return getShowcaseMapImage(this.map().mapImage)
  })

  readonly isAuthenticated = this.authService.isAuthenticatedSignal;

  onImgError(event: Event) {
    (event.target as HTMLImageElement).src = '/maps/map_placeholder.jpg';
  }

  onMapClick() {
    this.router.navigate(['/map', this.mapName()]);
  }

  onTagClick(tag: MapTag) {
    this.formService.form.categories().value.set(
      [tag.tagName]
    )
  }

}
