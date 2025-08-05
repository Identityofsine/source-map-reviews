import { ChangeDetectionStrategy, Component, computed, input, output } from '@angular/core';
import { IconComponent } from '@arch-shared/arch-ui';

@Component({
  selector: 'reviews-component',
  imports: [
    IconComponent
  ],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.scss'],
})
export class ReviewsComponent {

  readonly readOnly = input<boolean>(true);
  readonly reviewScore = input<number>(0);
  // out put of a "rating"
  readonly onStarClick = output<number>();


  readonly stars = computed(() =>
    Array.from({ length: 5 }, (_, i) => ({ idx: i, status: i < this.reviewScore() ? 'star' : 'no-star' }))
  );

}
