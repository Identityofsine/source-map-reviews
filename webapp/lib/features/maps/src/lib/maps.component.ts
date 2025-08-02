import { Component, inject, input } from '@angular/core';
import { rxResource } from '@angular/core/rxjs-interop';
import { MapsService } from '@arch-shared/data-source';
import { MapHeaderComponent } from './components/lib-map-header/lib-map-header.component';
import { MapGalleryComponent } from './components/lib-map-gallery/lib-map-gallery.component';
import { MapReviewsComponent } from './components/lib-map-reviews/lib-map-reviews.component';
import { ReviewsService } from 'lib/shared/data-source/src/lib/reviews.service';

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
  readonly reviewsService = inject(ReviewsService);

  readonly id = input.required<string>();

  private readonly _map = rxResource({
    request: () => ({ id: this.id() }),
    loader: ({ request }) => this.mapService.getMap(request.id),
  });

  private readonly _reviews = rxResource({
    request: () => ({ id: this.id() }),
    loader: ({ request }) => this.reviewsService.getReviews(request.id),
  });

  readonly map = this._map.value;
  readonly reviews = this._reviews.value;

}
