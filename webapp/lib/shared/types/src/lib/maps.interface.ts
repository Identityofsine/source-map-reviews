import { Image } from "./images.interface";

export interface MapGeneric<D> {
  mapName?: string;
  mapImage?: string;
  thumbnail?: Image;
  mapPath?: string;
  mapTags?: GenericMapTag<D>[];
}

export interface MapApi extends MapGeneric<string> {
}

export interface Map extends MapGeneric<Date> {
}

export interface GenericMapTag<D> {
  tagName: string;
  tagDescription?: string;
  tagDescriptionShort?: string;
  createdAt?: D
  updatedAt?: D;
}

export interface MapTagApi extends GenericMapTag<string> {
}

export interface MapTag extends GenericMapTag<Date> {
}

export interface MapSearchForm {
  searchTerm?: string;
  tags?: string[];
  reviewed?: boolean;
  unreviewed?: boolean;
}

export interface GenericTagLk<D> {
  tagLk: string;
  tagDescription?: string;
  tagDescriptionShort?: string;
  createdAt?: D;
  updatedAt?: D;
}

export interface TagLkApi extends GenericTagLk<string> {
}

export interface TagLk extends GenericTagLk<Date> {
}
