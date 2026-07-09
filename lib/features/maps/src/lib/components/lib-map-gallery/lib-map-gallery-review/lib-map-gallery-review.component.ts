import { DatePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapReview } from '@arch-shared/types';
import { UsernamePipe } from '@arch-shared/util';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-gallery-review',
  templateUrl: './lib-map-gallery-review.component.html',
  styleUrl: './lib-map-gallery-review.component.scss',
  imports: [DatePipe, UsernamePipe],
})
export class MapGalleryReviewComponent {

  readonly review = input<MapReview>();

}
