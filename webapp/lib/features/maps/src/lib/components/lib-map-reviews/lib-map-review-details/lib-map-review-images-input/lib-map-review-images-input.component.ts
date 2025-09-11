import { ChangeDetectionStrategy, Component } from '@angular/core';
import { MapReviewImagesInputThumbnail } from './components/lib-map-review-images-input-image-thumbnail.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review-images-input',
  templateUrl: './lib-map-review-images-input.component.html',
  styleUrls: ['./lib-map-review-images-input.component.scss'],
  imports: [MapReviewImagesInputThumbnail],
})
export class MapReviewImagesInputComponent {

}
