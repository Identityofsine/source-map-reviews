import { ChangeDetectionStrategy, Component, computed, HostBinding, input } from '@angular/core';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'arch-empty-result',
  template: `
  @if (items() && items().length === 0) {
    <div class="empty-result">
      <span>No results found.</span>
    </div>
  }
  `,
  styleUrl: './empty-result.component.scss',
  imports: [],
})
export class EmptyResultComponent {

  /**
   * items is treated as an input to the component.
   * when `undefined` or `null`, the component will not render anything.
   * when `[]`, it will render an empty result message.
   */
  readonly items = input<unknown[]>();

  readonly loading = input<boolean>(false);

  public readonly isEmpty = computed(() => {
    const items = this.items();

    if (this.loading()) return false; // Do not render if loading is true

    if (items === undefined || items === null) {
      return false; // Do not render if items is undefined or null
    }

    return items.length === 0; // Render if items is an empty array
  })

  @HostBinding('style.display')
  get style() {
    return this.isEmpty() ? 'block' : 'none';
  }



}
