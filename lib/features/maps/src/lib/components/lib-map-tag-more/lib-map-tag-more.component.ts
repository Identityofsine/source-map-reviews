import { ChangeDetectionStrategy, Component, computed, input, output } from '@angular/core';
import { MapTag } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-tags-remaining',
  template: `
    <button
      class="map-tag map-tag--remaining"
      (click)="onMoreClick($event)"
      type="button"
    >
      {{ displayText() }}
    </button>
  `,
  styleUrl: './lib-map-tag-more.component.scss',
})
export class MapTagsRemainingComponent {
  readonly remainingTags = input.required<MapTag[]>();
  readonly click = output<MapTag[]>();

  readonly displayText = computed(() => {
    const count = this.remainingTags().length;
    const displayCount = Math.min(count, 99);
    return `${displayCount}+`;
  });

  onMoreClick($event: MouseEvent): void {
    $event.stopPropagation();
    $event.preventDefault();
    this.click.emit(this.remainingTags());
  }
}
