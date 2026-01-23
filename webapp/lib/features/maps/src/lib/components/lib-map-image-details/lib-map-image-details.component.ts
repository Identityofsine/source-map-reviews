import { ChangeDetectionStrategy, Component, computed, effect, inject, output, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ArchFileUploadComponent, ArchTextAreaComponent, ArchModalComponent } from '@arch-shared/arch-ui';
import { Image } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-image-details',
  templateUrl: './lib-map-image-details.component.html',
  styleUrl: './lib-map-image-details.component.scss',
  imports: [
    ArchModalComponent,
    ArchTextAreaComponent,
    ArchFileUploadComponent,
    ReactiveFormsModule,
  ],
})
export class MapImageDetailsComponent {

  readonly fb = inject(FormBuilder);

  public shouldReloadImage = output<Image | false>();
  readonly form = this.fb.group({
    description: ['', [Validators.maxLength(500)]],
    fileName: ['', [Validators.required]],
  });

  readonly file = signal<File[] | null>(null);
  readonly fileName = toSignal(this.form.get('fileName').valueChanges)

  protected fileSelected(files: File[] | null) {
    this.file.set(files);
  }

  constructor() {
    effect(() => {
      const file = this.file();
      if (file && file.length > 0) {
        this.form.patchValue({ fileName: file?.[0].name });
      } else {
        this.form.patchValue({ fileName: '' });
      }
    })
  }

}
