import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapReview } from 'lib/shared/types/src/lib/reviews.interface';

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
