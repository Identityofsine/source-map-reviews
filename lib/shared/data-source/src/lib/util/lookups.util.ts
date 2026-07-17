import { LookupMap } from "lib/shared/types/src/lib/lookup.types";

export const composeLookups = (lookupsRaw: LookupMap): LookupMap => {
  const lookups = lookupsRaw.lookups;
  return {
    ...lookupsRaw,
    lookups: {
      ...lookups,
      webhookByLk: reduceIntoKeyedMap(lookups.webhookLk, (item) => item.webhookLk),
      gameByLk: reduceIntoKeyedMap(lookups.gameLk, (item) => item.lkGame),
      mapCategoryByLk: reduceIntoKeyedMap(lookups.mapCategoryLk, (item) => item.lkMapCategory),
    }
  }
}

const reduceIntoKeyedMap = <T extends Record<string, any>>(arr: T[], keySupplier: (item: T) => string): Record<string, T> => {
  return arr.reduce((acc, item) => {
    const key = keySupplier(item);
    acc[key] = item;
    return acc;
  }, {} as Record<string, T>);
}
