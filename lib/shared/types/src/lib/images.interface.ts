export interface Image {
  imageId?: number;
  imagePath?: string;
  caption?: string;
}

export interface MapImage {
  mapReviewId?: number;
  mapReviewImageId?: number;
  mapImageId?: number;
  imageId?: number;
  image?: Image;
}

export interface UploadImageResponse {
  images: Image[];
}

export interface ImageForm {
  json: string;
  files: File[];
}
