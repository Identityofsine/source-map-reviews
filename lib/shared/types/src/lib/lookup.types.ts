import { Lookups, Webhook } from "./gen.types";

interface BasicLk {
  description: string;
  shortDescription: string;
  createdAt: string;
  updatedAt: string;
}

export interface MapCategoryLk extends BasicLk {
  lkMapCategory: string;
}

export interface GameLk extends BasicLk {
  lkGame: string;
}

export type BaseLookups = {
  webhookLk: Webhook[];
  webhookByLk: Record<string, Webhook>;
  mapCategoryLk: MapCategoryLk[];
  mapCategoryByLk: Record<string, MapCategoryLk>;
  gameLk: GameLk[];
  gameByLk: Record<string, GameLk>;
}

export type LookupMap = {
  lookups: BaseLookups
} & Pick<Lookups, 'cacheTimestamp'>

export type LookupMapKeys = keyof LookupMap;

export type LookupKeys = keyof BaseLookups;

