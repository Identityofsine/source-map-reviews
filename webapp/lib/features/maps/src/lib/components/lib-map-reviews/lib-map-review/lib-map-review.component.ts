import { DatePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { ReviewsComponent } from '@arch-feature/reviews';
import { IconComponent } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review',
  templateUrl: './lib-map-review.component.html',
  styleUrls: ['./lib-map-review.component.scss'],
  imports: [
    IconComponent,
    DatePipe,
    ReviewsComponent,
  ],
})
export class MapReviewComponent {

  readonly review = input<MapReview>();
  readonly onReviewClick = output<MapReview>();

  readonly onReviewExpand = output<MapReview>();

}
