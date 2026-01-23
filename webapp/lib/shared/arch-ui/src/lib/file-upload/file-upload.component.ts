import { NgTemplateOutlet } from '@angular/common';
import { ChangeDetectionStrategy, Component, effect, input, output, signal, viewChild } from '@angular/core';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'file-upload',
  templateUrl: './file-upload.component.html',
  styleUrl: './file-upload.component.scss',
  imports: [NgTemplateOutlet],
})
export class ArchFileUploadComponent {

  readonly multiple = input<boolean>(false);
  readonly file = signal<File[] | null>(null);
  readonly fileSelected = output<File[] | null>();

  constructor() {
    effect(() => {
      const file = this.file();
      this.fileSelected.emit(file);
    })
  }

  onFileSelected(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      const filesArray = Array.from(input.files);
      this.file.set(filesArray);
    }
  }

}
