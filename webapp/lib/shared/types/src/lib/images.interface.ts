export interface Image {
  imageId?: string;
  imagePath?: string;
  caption?: string;
}

export interface MapImage {
  mapImageId?: string;
  mapName?: string;
  images?: Image[];
}
