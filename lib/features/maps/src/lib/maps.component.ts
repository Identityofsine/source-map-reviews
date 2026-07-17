import { Component, computed, inject, input, resource, signal } from '@angular/core';
import { rxResource } from '@angular/core/rxjs-interop';
import { LookupCacheService, MapsService } from '@arch-shared/data-source';
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

  readonly mapService = inject(MapsService);
  readonly lookupsService = inject(LookupCacheService)

  readonly id = input.required<string>();

  private readonly _map = rxResource({
    params: () => ({ id: this.id() }),
    stream: ({ params }) => this.mapService.getMap(params.id),
  });

  readonly categoryLks = this.lookupsService.getLookup('mapCategoryByLk');

  readonly map = this._map.value;
  readonly reviews = computed(() => this._map.value()?.mapReview ?? []);


  readonly categories = computed(() => {
    return this.map().categories;
  })

  readonly currentReview = signal<MapReview | null>(null);

  readonly mapImages = computed(() => this.reviews().filter(review => review.images) ?? []);

  protected onReviewClick(review: MapReview): void {
    if (review.images && review.images.length > 0) {
      this.currentReview.set(review);
    }
  }

  public reloadReviews(): void {
    this._map.reload();
  }

}
