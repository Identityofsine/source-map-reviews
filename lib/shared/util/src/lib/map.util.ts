import { MapImage } from "@arch-shared/types";

export function getShowcaseMapImage(mapImages: MapImage[]) {
  // sort with lowest seqNo first and return the first image
  return (mapImages || []).sort((a, b) => a.seqNo - b.seqNo)?.[0] as MapImage | undefined;
}
