import { ChangeDetectionStrategy, Component, forwardRef, input, signal, computed, inject, DestroyRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { DropdownSelectedItemComponent } from '../dropdown-selected-item/dropdown-selected-item.component';

export interface DropdownItem {
  key: string;
  value: string;
}

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'arch-dropdown-input',
  template: `
    @if (label()) {
      <label class="dropdown-input-label" [for]="inputId()">{{ label() }}</label>
    }

    <div class="dropdown-input-container">
      <!-- Selected items display -->
      @if (selectedItems().length > 0) {
        <div class="dropdown-selected-items">
          @for (item of selectedItems(); track item.key) {
            <arch-dropdown-selected-item
              [item]="item"
              (remove)="removeItem(item.key)"
            />
          }
        </div>
      }

      <!-- Text input -->
      <div class="dropdown-input-wrapper">
        <input
          #textInput
          class="dropdown-input"
          [id]="inputId()"
          [type]="'text'"
          [disabled]="_disabled()"
          [value]="currentInput()"
          [placeholder]="placeholder()"
          (input)="onInput($event)"
          (focus)="onFocus()"
          (blur)="onBlur()"
          (keydown)="onKeyDown($event)"
        />

        <!-- Dropdown toggle button -->
        <button
          type="button"
          class="dropdown-toggle"
          [disabled]="_disabled()"
          (click)="toggleDropdown()"
        >
          ▼
        </button>
      </div>

      <!-- Dropdown menu -->
      @if (showDropdown()) {
        <div class="dropdown-menu">
          @if (filteredItems().length === 0 && !freeRange()) {
            <div class="dropdown-no-items">No items available</div>
          }

          @for (item of filteredItems(); track item.key) {
            <div
              class="dropdown-item"
              [class.selected]="isSelected(item.key)"
              (click)="selectItem(item)"
            >
              {{ item.value }}
            </div>
          }

          @if (freeRange() && currentInput().trim() && !isExactMatch()) {
            <div
              class="dropdown-item dropdown-add-new"
              (click)="addCustomItem()"
            >
              Add "{{ currentInput().trim() }}"
            </div>
          }
        </div>
      }
    </div>
  `,
  styleUrls: ['./dropdown-input.component.scss'],
  imports: [DropdownSelectedItemComponent],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ArchDropdownInputComponent),
      multi: true,
    },
  ],
  standalone: true,
})
export class ArchDropdownInputComponent implements ControlValueAccessor {

  // Input properties
  readonly placeholder = input('');
  readonly disabled = input(false);
  readonly label = input<string | null>(null);
  readonly freeRange = input(false);
  readonly items = input<any[]>([]);
  readonly itemKey = input<string>('key');
  readonly itemValue = input<string>('value');
  readonly inputId = input(`dropdown-input-${Math.random().toString(36).substr(2, 9)}`);

  // Internal state
  protected readonly _disabled = signal(false);
  readonly mutatedItems = computed(() => {
    return (this.items() ?? [])?.map(item => ({
      key: item[this.itemKey()],
      value: item[this.itemValue()],
    })) as DropdownItem[];
  })
  readonly currentInput = signal('');
  readonly showDropdown = signal(false);
  readonly selectedKeys = signal<string[]>([]);

  // Computed properties
  readonly selectedItems = computed(() => {
    const keys = this.selectedKeys();
    const itemsMap = new Map(this.mutatedItems().map(item => [item.key, item]));

    return keys.map(key => {
      const existingItem = itemsMap.get(key);
      return existingItem || { key, value: key };
    });
  });

  readonly filteredItems = computed(() => {
    const input = this.currentInput().toLowerCase();
    const selected = new Set(this.selectedKeys());

    return this.mutatedItems()
      .filter(item =>
        !selected.has(item.key) &&
        item.value.toLowerCase().includes(input)
      );
  });

  readonly isExactMatch = computed(() => {
    const input = this.currentInput().toLowerCase();
    return this.mutatedItems().some(item => item.value.toLowerCase() === input);
  });

  // ControlValueAccessor implementation
  private onChange: (value: string[]) => void = () => { };
  private onTouched: () => void = () => { };

  writeValue(value: string[]): void {
    this.selectedKeys.set(value || []);
  }

  registerOnChange(fn: (value: string[]) => void): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void): void {
    this.onTouched = fn;
  }

  setDisabledState(isDisabled: boolean): void {
    this._disabled.set(isDisabled);
  }

  // Event handlers
  onInput(event: Event): void {
    const inputElement = event.target as HTMLInputElement;
    this.currentInput.set(inputElement.value);
    this.showDropdown.set(true);
  }

  onFocus(): void {
    this.showDropdown.set(true);
  }

  onBlur(): void {
    // Delay hiding dropdown to allow for item selection
    setTimeout(() => {
      this.showDropdown.set(false);
      this.onTouched();
    }, 150);
  }

  onKeyDown(event: KeyboardEvent): void {
    if (event.key === 'Enter' && this.freeRange() && this.currentInput().trim()) {
      event.preventDefault();
      this.addCustomItem();
    } else if (event.key === 'Escape') {
      this.showDropdown.set(false);
      this.currentInput.set('');
    }
  }

  toggleDropdown(): void {
    this.showDropdown.update(show => !show);
  }

  selectItem(item: DropdownItem): void {
    if (!this.isSelected(item.key)) {
      const newSelected = [...this.selectedKeys(), item.key];
      this.selectedKeys.set(newSelected);
      this.onChange(newSelected);
    }
    this.currentInput.set('');
    this.showDropdown.set(false);
  }

  removeItem(key: string): void {
    const newSelected = this.selectedKeys().filter(k => k !== key);
    this.selectedKeys.set(newSelected);
    this.onChange(newSelected);
  }

  addCustomItem(): void {
    if (!this.freeRange()) return;

    const input = this.currentInput().trim().toLowerCase();
    if (input && !this.isSelected(input)) {
      const newSelected = [...this.selectedKeys(), input];
      this.selectedKeys.set(newSelected);
      this.onChange(newSelected);
    }
    this.currentInput.set('');
    this.showDropdown.set(false);
  }

  isSelected(key: string): boolean {
    return this.selectedKeys().includes(key);
  }
}
