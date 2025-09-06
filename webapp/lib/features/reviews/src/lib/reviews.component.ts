import { ChangeDetectionStrategy, Component, computed, effect, input, output } from '@angular/core';
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
  readonly reviewScore = input<number>();
  // out put of a "rating"
  readonly onStarClick = output<number>();


  readonly stars = computed(() => {
    const reviewScore = this.reviewScore();
    if (!reviewScore || reviewScore === 0) return [];
    const array: { idx: number, status: 'star' | 'no-star' }[] = [];
    for (let i = 1; i <= 5; i++) {
      array.push(
        { idx: i - 1, status: i <= reviewScore ? 'star' : 'no-star' }
      )
    }
    return array;
  }, {
    equal: (a, b) => JSON.stringify(a) === JSON.stringify(b)
  });

  constructor() {
    console.log(this.stars())
  }

}
