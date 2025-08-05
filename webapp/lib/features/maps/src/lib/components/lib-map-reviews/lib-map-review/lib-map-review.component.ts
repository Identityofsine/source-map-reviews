import { DatePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { ReviewsComponent } from '@arch-feature/reviews';
import { MapReview } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review',
  templateUrl: './lib-map-review.component.html',
  styleUrls: ['./lib-map-review.component.scss'],
  imports: [
    DatePipe,
    ReviewsComponent,
  ],
})
export class MapReviewComponent {

  readonly review = input<MapReview>();

}
