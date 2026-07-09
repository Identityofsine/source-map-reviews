import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { IconComponent } from '../icon/icon.component';
import { DropdownItem } from '../dropdown-input/dropdown-input.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'arch-dropdown-selected-item',
  template: `
    <div class="dropdown-selected-item">
      <span class="dropdown-selected-item-text">{{ item().value }}</span>
      <button 
        type="button"
        class="dropdown-selected-item-remove"
        (click)="onRemove()"
        [title]="'Remove ' + item().value"
      >
        <arch-icon 
          class="remove-icon"
          [src]="'/ui/circle-x.svg'"
        />
      </button>
    </div>
  `,
  styleUrls: ['./dropdown-selected-item.component.scss'],
  imports: [IconComponent],
  standalone: true,
})
export class DropdownSelectedItemComponent {
  readonly item = input.required<DropdownItem>();
  readonly remove = output<string>();

  onRemove(): void {
    this.remove.emit(this.item().key);
  }
} 