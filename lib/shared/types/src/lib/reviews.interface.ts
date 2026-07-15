import { MapImage } from "./gen.types";

interface MapReviewGeneric<D> {
  mapReviewId?: number;
  mapName?: string;
  userId?: number;
  images?: MapImage[];
  stars?: number;
  review?: string;
  createdAt?: D;
  updatedAt?: D;
}

export interface MapReviewApi extends MapReviewGeneric<string> { }

