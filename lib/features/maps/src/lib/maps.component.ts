import { Component, computed, inject, input, signal } from '@angular/core';
import { rxResource } from '@angular/core/rxjs-interop';
import { MapsService, ReviewsService } from '@arch-shared/data-source';
import { MapHeaderComponent } from './components/lib-map-header/lib-map-header.component';
import { MapGalleryComponent } from './components/lib-map-gallery/lib-map-gallery.component';
import { MapReviewsComponent } from './components/lib-map-reviews/lib-map-reviews.component';
import { MapReview } from '@arch-shared/types';

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
    params: () => ({ id: this.id() }),
    stream: ({ params }) => this.mapService.getMap(params.id),
  });

  private readonly _reviews = rxResource({
    params: () => ({ id: this.id() }),
    stream: ({ params }) => this.reviewsService.getReviews(params.id),
  });


  readonly map = this._map.value;
  readonly reviews = this._reviews.value;

  readonly currentReview = signal<MapReview | null>(null);

  readonly mapImages = computed(() => this.reviews()?.filter(review => review.images) ?? []);

  protected onReviewClick(review: MapReview): void {
    if (review.images && review.images.length > 0) {
      this.currentReview.set(review);
    }
  }

  public reloadReviews(): void {
    this._reviews.reload();
  }

}
