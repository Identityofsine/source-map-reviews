import { ChangeDetectionStrategy, Component, computed, effect, inject, input, output } from '@angular/core';
import { ArchModalComponent, ArchTextInputComponent } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';
import { MapReviewDetailsFormService } from './lib-map-review-details-form.service';
import { CurrentUserService } from 'lib/shared/auth/src/lib/current-user.service';
import { ArchTextAreaComponent } from 'lib/shared/arch-ui/src/lib/text-area/text-area.component';
import { ReactiveFormsModule } from '@angular/forms';
import { MapReviewInputComponent } from './lib-map-review-input/lib-map-review-input.component';
import { ReviewsService } from 'lib/shared/data-source/src/lib/reviews.service';
import { toObservable, toSignal } from '@angular/core/rxjs-interop';
import { map, tap } from 'rxjs';
import { UsernamePipe } from '@arch-shared/util';
import { MapReviewImagesInputComponent } from './lib-map-review-images-input/lib-map-review-images-input.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review-details',
  templateUrl: './lib-map-review-details.component.html',
  styleUrls: ['./lib-map-review-details.component.scss'],
  imports: [
    ReactiveFormsModule,
    ArchModalComponent,
    ArchTextAreaComponent,
    UsernamePipe,
    MapReviewInputComponent,
    MapReviewImagesInputComponent,
  ],
})
export class MapReviewDetailsComponent {

  readonly formService = inject(MapReviewDetailsFormService);
  readonly currentUserService = inject(CurrentUserService)
  readonly reviewsService = inject(ReviewsService);

  readonly review = input<MapReview>();
  readonly mapName = input<string>();
  public readonly shouldReload = output<boolean>();
  readonly form = this.formService.form;

  readonly isAdding = computed(() => {
    return this.review() === undefined || this.review() === null;
  })

  readonly isEditing = computed(() => {
    return this.isAdding() || (this.review()?.userId === this.currentUserService.currentUser()?.id);
  });

  readonly isReadOnly = toSignal(this.formService.form.statusChanges.pipe(
    map(_ => this.form.disabled)
  ))

  constructor() {
    effect(() => {
      const curUser = this.currentUserService.currentUser();
      const isAdding = this.isAdding();
      this.formService.setReadOnly(!curUser || !this.isEditing());
      if (isAdding) {
        if (!curUser) {
          throw new Error('User must be logged in to add a review');
        }
        this.formService.populateFormWithEmptyReview(
          curUser?.id,
          this.mapName()
        );
      } else {
        this.formService.populateFormWithMapReview(
          this.review()
        );
      }
    })
  }

  onSubmit() {
    this.reviewsService.saveReview(this.form.value)
      .pipe().subscribe((res) => {
        if (res.mapReviewId) {
          this.shouldReload.emit(true);
        }
      })
  }

}
