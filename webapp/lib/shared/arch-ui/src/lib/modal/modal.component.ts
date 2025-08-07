import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, input, linkedSignal, OnInit, output, Signal, TemplateRef, viewChild } from '@angular/core';
import { OverlayModule } from '@angular/cdk/overlay';
import { IconComponent } from '../icon/icon.component';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'arch-modal',
  template: `
  <ng-template
  cdkConnectedOverlay
  [cdkConnectedOverlayOrigin]="modal()"
  [cdkConnectedOverlayHasBackdrop]="true"
  [cdkConnectedOverlayOpen]="isOpen()"
  >
    <div modal class="arch-modal">
      <div class="arch-modal-header">
        <div modal-title class="arch-modal-title">
          <ng-content select="[modal-title]">Mimi</ng-content>
        </div>
        <button modal-close class="arch-modal-close" (click)="closeModal()">

        </button>


      </div>
      <ng-content></ng-content>
    </div>
  </ng-template>
  `,
  styleUrl: './modal.component.scss',
  imports: [CommonModule, OverlayModule, IconComponent],
})
export class ArchModalComponent {

  readonly contentTemplate: Signal<TemplateRef<unknown>> = viewChild('content');
  readonly modal: Signal<TemplateRef<unknown>> = viewChild('modal');
  readonly cdkConnectedOverlayOrigin: Signal<TemplateRef<unknown>> = viewChild('cdkConnectedOverlayOrigin');

  readonly open = input<boolean>(false);
  readonly closed = output<void>();

  readonly isOpen = linkedSignal(() => this.open());

  closeModal() {
    this.isOpen.set(false);
    this.closed.emit();
  }

}
