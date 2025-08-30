import { ChangeDetectionStrategy, Component, computed, effect, inject, input } from '@angular/core';
import { ArchModalComponent, ArchTextInputComponent } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';
import { MapReviewDetailsFormService } from './lib-map-review-details-form.service';
import { CurrentUserService } from 'lib/shared/auth/src/lib/current-user.service';
import { ArchTextAreaComponent } from 'lib/shared/arch-ui/src/lib/text-area/text-area.component';
import { ReactiveFormsModule } from '@angular/forms';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review-details',
  templateUrl: './lib-map-review-details.component.html',
  styleUrls: ['./lib-map-review-details.component.scss'],
  imports: [
    ReactiveFormsModule,
    ArchModalComponent,
    ArchTextAreaComponent,
  ],
})
export class MapReviewDetailsComponent {


  readonly formService = inject(MapReviewDetailsFormService);
  readonly currentUserService = inject(CurrentUserService)

  readonly review = input<MapReview>();
  readonly mapName = input<string>();
  readonly form = this.formService.form;

  readonly isAdding = computed(() => {
    return this.review() === undefined || this.review() === null;
  })

  readonly isEditing = computed(() => {
    return true;
  });

  constructor() {
    effect(() => {
      const curUser = this.currentUserService.currentUser();
      const isAdding = this.isAdding();
      this.formService.setReadOnly(this.isEditing());
      if (isAdding) {
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

}
