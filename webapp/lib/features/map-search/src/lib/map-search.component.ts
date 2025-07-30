import { Component, inject } from '@angular/core';
import { rxResource, toSignal } from '@angular/core/rxjs-interop';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { ArchTextInputComponent } from '@arch-shared/arch-ui';
import { MapsService } from '@arch-shared/data-source';
import { MapThumbnailComponent } from './components/lib-map-thumbnail/lib-map-thumbnail.component';

@Component({
  selector: 'arch-map-search',
  imports: [
    ArchTextInputComponent,
    ReactiveFormsModule,
    MapThumbnailComponent,
  ],
  templateUrl: './map-search.component.html',
  styleUrl: './map-search.component.scss',
})
export class MapSearchComponent {

  readonly fb = inject(FormBuilder);
  readonly mapService = inject(MapsService);

  readonly form = this.fb.group({
    searchTerm: ['']
  })

  readonly searchTerm = toSignal(this.form?.get('searchTerm')!.valueChanges, { initialValue: '' });

  readonly search = rxResource({
    request: () => ({
      searchTerm: this.searchTerm()
    }),
    loader: () => this.mapService.searchMaps({
      searchTerm: this.searchTerm() ?? '',
    }),
  });

}
