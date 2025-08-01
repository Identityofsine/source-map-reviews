import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { ArchContainer } from '@arch-shared/arch-ui';
import { MapReview } from 'lib/shared/types/src/lib/reviews.interface';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-gallery',
  templateUrl: './lib-map-gallery.component.html',
  styleUrl: './lib-map-gallery.component.scss',
  imports: [
    ArchContainer
  ],
})
export class MapGalleryComponent {

  readonly mapImages = input<MapReview[]>();

}
