import { ChangeDetectionStrategy, Component, output } from '@angular/core';
import { ArchModalComponent } from '@arch-shared/arch-ui';
import { Image } from '@arch-shared/types';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-image-details',
  templateUrl: './lib-map-image-details.component.html',
  styleUrl: './lib-map-image-details.component.scss',
  imports: [
    ArchModalComponent,
  ],
})
export class MapImageDetailsComponent {

  public shouldReloadImage = output<Image | false>();

}
