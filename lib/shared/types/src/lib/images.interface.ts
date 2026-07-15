export interface Image {
  imageId?: number;
  imagePath?: string;
  caption?: string;
}

export interface UploadImageResponse {
  images: Image[];
}

export interface ImageForm {
  json: string;
  files: File[];
}
