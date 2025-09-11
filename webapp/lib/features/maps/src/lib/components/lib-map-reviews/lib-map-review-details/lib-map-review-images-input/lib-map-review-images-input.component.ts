import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapReviewImagesInputThumbnail } from './components/lib-map-review-images-input-image-thumbnail.component';
import { AddButtonComponent } from '../../../lib-add-button/lib-add-button.component';
import { MapImage } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review-images-input',
  templateUrl: './lib-map-review-images-input.component.html',
  styleUrls: ['./lib-map-review-images-input.component.scss'],
  imports: [MapReviewImagesInputThumbnail, AddButtonComponent],
})
export class MapReviewImagesInputComponent {
  readonly images = input<MapImage[]>([]);


  public openAddImage() {
    const view = this.containerRef.createComponent(MapImageDetailsComponent);
    //view.setInput();
    view.instance.shouldReloadImage.subscribe((res) => {
      if (res) {
        this.reloadImages();
      }
      view.destroy();
    });
  }
}
