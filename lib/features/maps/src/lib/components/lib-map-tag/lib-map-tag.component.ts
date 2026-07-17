import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MapTag } from '@arch-shared/types';
import { MapCategoryLk } from 'lib/shared/types/src/lib/lookup.types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-tag',
  templateUrl: './lib-map-tag.component.html',
  styleUrls: ['./lib-map-tag.component.scss'],
  imports: [],
})
export class MapTagComponent {

  readonly category = input<MapCategoryLk>();
  readonly categoryClick = output<MapCategoryLk>();

  onCategoryClick($event: MouseEvent) {
    $event.preventDefault();
    $event.stopPropagation();
    this.categoryClick.emit(this.category());
  }

}
