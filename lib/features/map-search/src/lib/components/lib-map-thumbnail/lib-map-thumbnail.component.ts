import { ChangeDetectionStrategy, Component, computed, inject, input } from '@angular/core';
import { Map, MapTag } from '@arch-shared/types';
import { Router } from '@angular/router';
import { MapTagsComponent } from '@arch-feature/maps';
import { AuthService } from '@arch-shared/auth';
import { MapSearchFormService } from '../../map-search-form.service';

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

  readonly map = input<Map>();

  readonly mapName = computed(() => {
    return this.map()?.mapName || '';
  });

  readonly mapTags = computed(() => {
    return this.map().mapTags || [];
  })

  readonly mapImage = computed(() => {
    return this.map()?.thumbnail.imagePath || ''
  })

  readonly isAuthenticated = this.authService.isAuthenticatedSignal;

  onImgError(event: Event) {
    (event.target as HTMLImageElement).src = '/maps/map_placeholder.jpg';
  }

  onMapClick() {
    this.router.navigate(['/map', this.mapName()]);
  }

  onTagClick(tag: MapTag) {
    this.formService.form.get('tags')?.setValue(
      [tag.tagName]
    )
  }

}
