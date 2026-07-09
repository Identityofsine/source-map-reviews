import { ChangeDetectionStrategy, Component, computed, inject, input, output, ViewContainerRef } from '@angular/core';
import { ArchContainer } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';
import { AddButtonComponent } from '../lib-add-button/lib-add-button.component';
import { AuthService } from '@arch-shared/auth';
import { MapReviewComponent } from './lib-map-review/lib-map-review.component';
import { MapReviewDetailsComponent } from './lib-map-review-details/lib-map-review-details.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-reviews',
  templateUrl: './lib-map-reviews.component.html',
  styleUrls: ['./lib-map-reviews.component.scss'],
  imports: [
    ArchContainer,
    MapReviewComponent,
    AddButtonComponent,
  ]
})
export class MapReviewsComponent {

  readonly authService = inject(AuthService);
  readonly viewContainer = inject(ViewContainerRef);

  readonly reviews = input<MapReview[]>();
  readonly mapName = input<string>();
  readonly onReviewClick = output<MapReview>();
  readonly shouldReloadReviews = output<void>();

  readonly isEmpty = computed(() => {
    return (this.reviews() ?? []).length <= 0
  })

  readonly shouldShow = this.authService.isAuthenticatedSignal;


  onAddClick() {
    const containerRef = this.viewContainer.createComponent(MapReviewDetailsComponent)
    containerRef.setInput('review', undefined);
    containerRef.setInput('mapName', this.mapName()); // Set the map name if needed
    containerRef.instance.shouldReload.subscribe(() => {
      containerRef.destroy();
      this.shouldReloadReviews.emit();
    })
  }

  onReviewClickHandler(review: MapReview) {
    this.onReviewClick.emit(review);
  }

  openAlreadyReviewModal(review: MapReview) {
    const containerRef = this.viewContainer.createComponent(MapReviewDetailsComponent);
    containerRef.setInput('review', review);
    containerRef.setInput('mapName', this.mapName());
    containerRef.instance.shouldReload.subscribe(() => {
      containerRef.destroy();
      this.shouldReloadReviews.emit();
    })
  }

}
