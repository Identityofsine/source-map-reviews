import { Lookups, Webhook } from "./gen.types";

export type LookupMap = {
  lookups: {
    webhookLk: Webhook[];
  }
} & Pick<Lookups, 'cacheTimestamp'>

export type LookupMapKeys = keyof LookupMap;
