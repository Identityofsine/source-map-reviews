import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapReview } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-gallery-review',
  templateUrl: './lib-map-gallery-review.component.html',
  styleUrl: './lib-map-gallery-review.component.scss',
  imports: [],
})
export class MapGalleryReviewComponent {

  readonly review = input<MapReview>();

}
