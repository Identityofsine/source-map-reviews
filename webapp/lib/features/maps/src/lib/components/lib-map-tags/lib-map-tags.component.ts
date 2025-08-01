import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MapTag } from '@arch-shared/types';
import { MapTagComponent } from '../lib-map-tag/lib-map-tag.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-tags',
  templateUrl: './lib-map-tags.component.html',
  styleUrl: './lib-map-tags.component.scss',
  imports: [
    MapTagComponent
  ],
})
export class MapTagsComponent {

  readonly tags = input<MapTag[]>();

}
