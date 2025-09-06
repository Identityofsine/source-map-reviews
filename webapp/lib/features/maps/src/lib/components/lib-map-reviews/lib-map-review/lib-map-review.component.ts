import { DatePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { ReviewsComponent } from '@arch-feature/reviews';
import { IconComponent } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';
import { UsernamePipe } from '@arch-shared/util';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review',
  templateUrl: './lib-map-review.component.html',
  styleUrls: ['./lib-map-review.component.scss'],
  imports: [
    IconComponent,
    UsernamePipe,
    DatePipe,
    ReviewsComponent,
  ],
})
export class MapReviewComponent {

  readonly review = input<MapReview>();
  readonly onReviewClick = output<MapReview>();

  readonly onReviewExpand = output<MapReview>();

}
