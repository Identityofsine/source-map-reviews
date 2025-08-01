import { Component, inject, input } from '@angular/core';
import { rxResource } from '@angular/core/rxjs-interop';
import { MapsService } from '@arch-shared/data-source';
import { MapHeaderComponent } from './components/lib-map-header/lib-map-header.component';

@Component({
  selector: 'arch-maps',
  imports: [
    MapHeaderComponent,
  ],
  templateUrl: './maps.component.html',
  styleUrl: './maps.component.scss',
})
export class MapsComponent {

  //DI
  readonly mapService = inject(MapsService);

  readonly id = input.required<string>();

  private readonly _map = rxResource({
    request: () => ({ id: this.id() }),
    loader: ({ request }) => this.mapService.getMap(request.id),
  });

  readonly map = this._map.value;

}
