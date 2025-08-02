import { DatePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapReview } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review',
  templateUrl: './lib-map-review.component.html',
  styleUrls: ['./lib-map-review.component.scss'],
  imports: [DatePipe],
})
export class MapReviewComponent {

  readonly review = input<MapReview>();

}
