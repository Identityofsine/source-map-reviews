import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MapImage } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review-images-input-thumbnail',
  templateUrl: './lib-map-review-images-input-image-thumbnail.component.html',
  styleUrl: 'lib-map-review-images-input-image-thumbnail.component.scss',
  imports: [],
})
export class MapReviewImagesInputThumbnail {
  readonly mapImage = input<MapImage>();
  readonly click = output<MapImage>();
}
