import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'arch-icon',
  template: `
    <div
    [className]="_class()"
    [style.--url]="_src()"
    >
    </div>
  `,
  styleUrl: './icon.component.scss',
  imports: [],
})
export class IconComponent {

  readonly class = input<string>('add-icon');
  readonly src = input<string>();

  readonly _src = computed(() =>
    `url(${this.src()})`)

  readonly _class = computed(() => {
    return this.class() ? `icon ${this.class()}` : 'icon';
  });

}
