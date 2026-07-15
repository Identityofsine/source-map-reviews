import { ChangeDetectionStrategy, Component, inject, input, output, ViewContainerRef } from '@angular/core';
import { MapReviewImagesInputThumbnail } from './components/lib-map-review-images-input-image-thumbnail.component';
import { AddButtonComponent } from '../../../lib-add-button/lib-add-button.component';
import { MapImage } from '@arch-shared/types';
import { MapImageDetailsComponent } from '../../../lib-map-image-details/lib-map-image-details.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review-images-input',
  templateUrl: './lib-map-review-images-input.component.html',
  styleUrls: ['./lib-map-review-images-input.component.scss'],
  imports: [MapReviewImagesInputThumbnail, AddButtonComponent],
})
export class MapReviewImagesInputComponent {

  readonly containerRef = inject(ViewContainerRef);
  readonly images = input<MapImage[]>([]);
  readonly imageAdded = output<MapImage>();


  public openAddImage() {
    const view = this.containerRef.createComponent(MapImageDetailsComponent);
    //view.setInput();
    view.instance.shouldReloadImage.subscribe((res) => {
      if (res) {
      }
      view.destroy();
    });
  }
}
