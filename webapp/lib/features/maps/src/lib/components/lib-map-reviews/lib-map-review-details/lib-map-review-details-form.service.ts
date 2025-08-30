import { inject, Injectable } from "@angular/core";
import { MapReviewsComponent } from "../lib-map-reviews.component";
import { FormBuilder } from "@angular/forms";
import { MapImage, MapReview } from "@arch-shared/types";

@Injectable({
  providedIn: 'root',
})
export class MapReviewDetailsFormService {

  readonly fb = inject(FormBuilder);

  readonly form = this.fb.group({
    userId: [0],
    mapName: [''],
    review: [''],
    stars: [0],
    images: [[] as MapImage[]],
  })

  public populateFormWithMapReview(review: MapReview) {
    this.form.patchValue({
      userId: review.userId,
      mapName: review.mapName,
      review: review.review,
      stars: review.stars,
      images: review.images,
    });
  }

  public populateFormWithEmptyReview(userId: number, mapName: string) {
    console.log('Populating form with empty review', userId, mapName);
    this.form.patchValue({
      userId: userId,
      mapName: mapName,
      review: '',
      stars: 0,
      images: [],
    });
  }

  public setReadOnly(readOnly: boolean) {
    if (readOnly) {
      this.form.disable();
    } else {
      this.form.enable();
    }
  }

}
