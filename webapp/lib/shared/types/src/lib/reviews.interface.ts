import { MapImage } from "./images.interface";

interface MapReviewGeneric<D> {
  mapReviewId?: string;
  mapName?: string;
  userId?: number;
  images?: MapImage[];
  stars?: number;
  review?: string;
  createdAt?: D;
  updatedAt?: D;
}

export interface MapReviewApi extends MapReviewGeneric<string> { }

export interface MapReview extends MapReviewGeneric<Date> { }

