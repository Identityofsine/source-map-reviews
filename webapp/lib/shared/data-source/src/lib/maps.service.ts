import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Map, MapApi, MapSearchForm } from '../../../types/src';
import { map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MapsService {

  readonly http = inject(HttpClient);
  readonly API_URL = `/api/maps`

  public getMaps(): Observable<Map[]> {
    return this.http.get<MapApi[]>(this.API_URL).pipe(
      map(maps => maps.map(map => this.populateMapFromBackend(map)))
    )
  }

  public searchMaps(form: MapSearchForm): Observable<Map[]> {
    return this.http.post<MapApi[]>(`${this.API_URL}/search`, form).pipe(
      map(maps => maps.map(map => this.populateMapFromBackend(map)))
    )
  }


  private populateMapFromBackend(map: MapApi): Map {
    return {
      ...map,
      mapTags: map?.mapTags?.map(tag => ({
        ...tag,
        createdAt: tag.createdAt ? new Date(tag.createdAt) : undefined,
        updatedAt: tag.updatedAt ? new Date(tag.updatedAt) : undefined
      }))
    }
  }

}
