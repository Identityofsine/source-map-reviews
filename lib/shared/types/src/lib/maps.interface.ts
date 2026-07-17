import { Map } from "./gen.types";
import { Image } from "./images.interface";
import { MapCategoryLk } from "./lookup.types";

export interface MapGeneric<D> {
  mapName?: string;
  mapImage?: string;
  thumbnail?: Image;
  mapPath?: string;
  mapTags?: GenericMapTag<D>[];
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

export type MapApi = {
  categories?: MapCategoryLk[];
} & Omit<Map, 'categories'>
