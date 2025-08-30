import { ChangeDetectionStrategy, Component, forwardRef, input, linkedSignal } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'arch-text-area',
  template: `
    @if (label()) {
      <label class="text-input-label" [for]="label()">{{ label() }}</label>
    }
    <textarea
      class="text-input"
      type="text"
      [disabled]="_disabled()"
      [value]="value"
      [placeholder]="placeholder()"
      (input)="onInput($event)"
    ></textarea>
  `,
  styleUrls: ['./text-area.component.scss'],
  imports: [

  ],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ArchTextAreaComponent),
      multi: true,
    },
  ],
  standalone: true,
})
export class ArchTextAreaComponent implements ControlValueAccessor {

  // This component can be extended in the future to include additional functionality

  //form stuff
  readonly placeholder = input('');
  readonly disabled = input(false);
  readonly readOnly = input(false);
  readonly label = input<string | null>(null);
  readonly formControlName = input<string | null>(null);

  protected readonly _disabled = linkedSignal(() => this.disabled());
  value = '';

  onChange: (value: string) => void = () => { };
  onTouched: () => void = () => { };

  writeValue(obj: string): void {
    this.value = obj || '';
  }

  registerOnChange(fn: any): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  setDisabledState(isDisabled: boolean): void {
    this._disabled.set(isDisabled);
  }

  onInput(event: Event): void {
    if (this.readOnly()) {
      return; // Do not allow input if readOnly is true
    }
    const inputElement = event.target as HTMLInputElement;
    this.value = inputElement.value;
    this.onChange(this.value);
  }
}
