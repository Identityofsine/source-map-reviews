import { ChangeDetectionStrategy, Component, computed, DestroyRef, effect, forwardRef, inject, Injector, input, linkedSignal, OnInit } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { ControlValueAccessor, NG_VALUE_ACCESSOR, NgControl, } from '@angular/forms';
import { IconComponent } from '@arch-shared/arch-ui';
import { BehaviorSubject, tap } from 'rxjs';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-review-input',
  templateUrl: './lib-map-review-input.component.html',
  styleUrl: './lib-map-review-input.component.scss',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => MapReviewInputComponent),
      multi: true,
    }
  ],
  imports: [
    IconComponent
  ],
  standalone: true,
})
export class MapReviewInputComponent implements ControlValueAccessor, OnInit {


  readonly formControlName = input<string | null>(null);
  readonly disabled = input(false);
  readonly readOnly = input(false);
  readonly range = Array(5).fill(0).map((_, i) => i + 1); // [1, 2, 3, 4, 5]
  readonly valueInput = input<number | null>(null);
  protected ngControl: NgControl;
  readonly injector = inject(Injector);
  readonly destroyRef = inject(DestroyRef);

  readonly value$ = new BehaviorSubject<number>(0);
  readonly valueSignal = toSignal(this.value$);
  value = 0;

  readonly stars = computed(() => {
    const reviewScore = this.valueSignal() ?? 0;
    const array: { idx: number, status: 'star' | 'no-star' }[] = [];
    for (let i = 1; i <= 5; i++) {
      array.push(
        { idx: i - 1, status: i <= reviewScore ? 'star' : 'no-star' }
      )
    }
    return array;
  }, {
    equal: (a, b) => JSON.stringify(a) === JSON.stringify(b)
  });



  protected readonly _disabled = linkedSignal(() => this.disabled());

  constructor() {
    effect(() => {
      const val = this.valueSignal();
      this.value = val;
      this.onChange(this.value);
      this.onTouched();
    })

  }

  ngOnInit() {
    this.ngControl = this.injector.get(NgControl);
    if (this.ngControl != null) {
      this.ngControl.valueChanges?.pipe(
        tap(value => this.value$.next(value || 0)),
      ).subscribe();
    }
  }

  onChange: (value: number) => void = () => { };
  onTouched: () => void = () => { };

  writeValue(obj: number): void {
    this.value$.next(obj);
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

  setStars(stars: number) {
    if (this.readOnly()) {
      return;
    }
    this.value$.next(stars);
  }

}
