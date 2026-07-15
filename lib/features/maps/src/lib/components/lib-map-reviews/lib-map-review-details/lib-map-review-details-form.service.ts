import { inject, Injectable } from "@angular/core";
import { MapReviewsComponent } from "../lib-map-reviews.component";
import { FormBuilder, Validators } from "@angular/forms";
import { MapImage, MapReview } from "@arch-shared/types";

@Injectable({
  providedIn: 'root',
})
export class MapReviewDetailsFormService {

  readonly fb = inject(FormBuilder);

  readonly form = this.fb.group({
    mapReviewId: [0],
    userId: [0],
    mapName: ['', [Validators.required]],
    review: ['', [Validators.required]],
    stars: [0, [Validators.required, Validators.min(1), Validators.max(5)]],
    images: [[] as MapImage[]],
  })

  public populateFormWithMapReview(review: MapReview) {
    this.form.patchValue({
    });
  }

  public populateFormWithEmptyReview(userId: number, mapName: string) {
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
