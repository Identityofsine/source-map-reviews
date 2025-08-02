import { HttpClient } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { MapReview, MapReviewApi } from "@arch-shared/types";
import { map, Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class ReviewsService {

  readonly http = inject(HttpClient);
  readonly API_URL = `/api/reviews`

  getReviews(mapName: string): Observable<MapReview[]> {
    return this.http.get<MapReviewApi[]>(`${this.API_URL}/${mapName}`).pipe(
      map(reviews => reviews.map(review => this.populateReviewFromBackend(review)))
    );
  }

  private populateReviewFromBackend(review: MapReviewApi): MapReview {
    return {
      ...review,
      createdAt: new Date(review.createdAt),
      updatedAt: new Date(review.updatedAt)
    };
  }


}
