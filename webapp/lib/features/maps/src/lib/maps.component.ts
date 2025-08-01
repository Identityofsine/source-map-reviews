import { Component, inject, input } from '@angular/core';
import { rxResource } from '@angular/core/rxjs-interop';
import { MapsService } from '@arch-shared/data-source';

@Component({
  selector: 'arch-maps',
  imports: [],
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
