import { ChangeDetectionStrategy, Component, computed, HostBinding, input, linkedSignal, output, viewChild, effect, ElementRef, afterNextRender, signal, OnDestroy } from '@angular/core';
import { MapTag } from '@arch-shared/types';
import { MapTagComponent } from '../lib-map-tag/lib-map-tag.component';
import { AddButtonComponent } from '../lib-add-button/lib-add-button.component';
import { MapTagsRemainingComponent } from '../lib-map-tag-more/lib-map-tag-more.component';
import { MapCategoryLk } from 'lib/shared/types/src/lib/lookup.types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-tags',
  templateUrl: './lib-map-tags.component.html',
  styleUrl: './lib-map-tags.component.scss',
  imports: [
    MapTagComponent,
    AddButtonComponent,
    MapTagsRemainingComponent
  ],
})
export class MapTagsComponent implements OnDestroy {

  readonly categories = input<MapCategoryLk[]>();
  readonly shouldShowAddButton = input<boolean>(false);
  readonly tagClick = output<MapTag>();
  readonly remainingTagsClick = output<MapTag[]>();

  readonly container = viewChild<ElementRef<HTMLElement>>('container');

  readonly isEmpty = computed(() => {
    return !this.categories() || this.categories()!.length === 0;
  });

  readonly visibleTagsCount = signal(0);

  readonly visibleTags = computed(() => {
    const tags = this.categories() || [];
    const count = this.visibleTagsCount();
    return tags.slice(0, count);
  });

  readonly remainingTags = computed(() => {
    const tags = this.categories() || [];
    const count = this.visibleTagsCount();
    return tags.slice(count);
  });

  readonly hasRemainingTags = computed(() => {
    return this.remainingTags().length > 0;
  });

  private resizeObserver?: ResizeObserver;
  private isAdjusting = false;

  constructor() {
    afterNextRender(() => {
      this.setupResizeObserver();
      this.adjustVisibleTags();
    });

    effect(() => {
      const tags = this.categories();
      if (tags) {
        this.visibleTagsCount.set(tags.length);
        setTimeout(() => this.adjustVisibleTags(), 50);
      }
    });
  }

  ngOnDestroy(): void {
    this.resizeObserver?.disconnect();
  }

  private setupResizeObserver(): void {
    const containerEl = this.container()?.nativeElement;
    if (!containerEl || typeof ResizeObserver === 'undefined') return;

    this.resizeObserver = new ResizeObserver(() => {
      if (!this.isAdjusting) {
        setTimeout(() => this.adjustVisibleTags(), 50);
      }
    });

    this.resizeObserver.observe(containerEl);
  }

  private async adjustVisibleTags(): Promise<void> {
    const containerEl = this.container()?.nativeElement;
    const tags = this.categories();

    if (!containerEl || !tags?.length || this.isAdjusting) return;

    this.isAdjusting = true;

    try {
      // Start with all tags
      let count = tags.length;
      this.visibleTagsCount.set(count);

      // Wait for initial DOM update
      await this.waitForDOMUpdate();

      // If no overflow with all tags, we're done
      if (!this.hasOverflow(containerEl)) {
        return;
      }

      // Reduce one by one with proper timing
      while (count > 0 && this.hasOverflow(containerEl)) {
        count--;
        this.visibleTagsCount.set(count);
        await this.waitForDOMUpdate();
      }

    } finally {
      this.isAdjusting = false;
    }
  }

  private waitForDOMUpdate(): Promise<void> {
    return new Promise(resolve => {
      requestAnimationFrame(() => {
        requestAnimationFrame(() => {
          resolve();
        });
      });
    });
  }

  private hasOverflow(container: HTMLElement): boolean {
    return container.scrollWidth > container.clientWidth + 1;
  }

  onRemainingTagsClick(remainingTags: MapTag[]): void {
    this.remainingTagsClick.emit(remainingTags);
  }

  onTagClick(event: any, tag: MapTag): void {
    this.tagClick.emit(tag);
  }

  onAddButtonClick(event: any): void {
    // Handle add button click if needed
  }

}
