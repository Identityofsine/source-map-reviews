import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapTag } from '@arch-shared/types';
import { MapTagsComponent } from '../lib-map-tags/lib-map-tags.component';
import { MapReviewsComponent } from '../lib-map-reviews/lib-map-reviews.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-header',
  templateUrl: './lib-map-header.component.html',
  styleUrls: ['./lib-map-header.component.scss'],
  imports: [
    MapTagsComponent,
  ],
})
export class MapHeaderComponent {

  readonly mapName = input<string>();
  readonly mapTags = input<MapTag[]>();

}
