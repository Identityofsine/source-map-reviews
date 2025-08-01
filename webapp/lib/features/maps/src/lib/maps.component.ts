import { Component, inject, input } from '@angular/core';
import { rxResource } from '@angular/core/rxjs-interop';
import { MapsService } from '@arch-shared/data-source';
import { MapHeaderComponent } from './components/lib-map-header/lib-map-header.component';
import { MapGalleryComponent } from './components/lib-map-gallery/lib-map-gallery.component';
import { MapReviewsComponent } from './components/lib-map-reviews/lib-map-reviews.component';

@Component({
  selector: 'arch-maps',
  imports: [
    MapHeaderComponent,
    MapReviewsComponent,
    MapGalleryComponent,
  ],
  templateUrl: './maps.component.html',
  styleUrl: './maps.component.scss',
})
export class MapsComponent {

  //DI
  readonly mapService = inject(MapsService);

  readonly id = input.required<string>();

  private readonly _map = rxResource({
    request: () => ({ id: this.id() }),
    loader: ({ request }) => this.mapService.getMap(request.id),
  });

  readonly map = this._map.value;

}
