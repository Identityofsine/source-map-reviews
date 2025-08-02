import { ChangeDetectionStrategy, Component, computed, inject, input } from '@angular/core';
import { ArchContainer } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';
import { AddButtonComponent } from '../lib-add-button/lib-add-button.component';
import { AuthService } from '@arch-shared/auth';
import { MapReviewComponent } from './lib-map-review/lib-map-review.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-reviews',
  templateUrl: './lib-map-reviews.component.html',
  styleUrls: ['./lib-map-reviews.component.scss'],
  imports: [
    ArchContainer,
    MapReviewComponent,
    AddButtonComponent,
  ],
})
export class MapReviewsComponent {

  readonly authService = inject(AuthService);

  readonly reviews = input<MapReview[]>();

  readonly isEmpty = computed(() => {
    return (this.reviews() ?? []).length <= 0
  })

  readonly shouldShow = this.authService.isAuthenticatedSignal;

}
