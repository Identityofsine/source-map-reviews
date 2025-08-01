import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';
import { ArchContainer } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-reviews',
  templateUrl: './lib-map-reviews.component.html',
  styleUrls: ['./lib-map-reviews.component.scss'],
  imports: [
    ArchContainer
  ],
})
export class MapReviewsComponent {

  readonly reviews = input<MapReview[]>();

  readonly isEmpty = computed(() => {
    return (this.reviews() ?? []).length <= 0
  })

}
