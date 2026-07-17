import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Map, MapApi, MapSearchForm, TagLk } from '../../../types/src';
import { catchError, filter, from, map, Observable, OperatorFunction, skipUntil, switchMap, take } from 'rxjs';
import { LookupCacheService } from './lookups.service';
import { toObservable } from '@angular/core/rxjs-interop';

@Injectable({
  providedIn: 'root'
})
export class MapsService {

  readonly http = inject(HttpClient);
  readonly #lookups = inject(LookupCacheService);
  readonly API_URL = `/api/maps`

  readonly loading$ = toObservable(this.#lookups.loading);

  readonly categoryLks = this.#lookups.getLookup('mapCategoryByLk');

  public getMaps(): Observable<MapApi[]> {
    return this.http.get<Map[]>(this.API_URL).pipe(
      map(maps => maps.map(map => this.populateMapFromBackend(map)))
    )
  }

  public getMap(id: string): Observable<MapApi> {
    return this.http.get<Map>(`${this.API_URL}/${id}`).pipe(
      map(map => this.populateMapFromBackend(map))
    )
  }

  public getTags(): Observable<TagLk[]> {
    return this.http.get<TagLk[]>(`${this.API_URL}/tags`).pipe(
      map(tags => tags.map(tag => ({
        ...tag,
        createdAt: tag.createdAt ? new Date(tag.createdAt) : undefined,
        updatedAt: tag.updatedAt ? new Date(tag.updatedAt) : undefined
      })))
    )
  }

  public searchMaps(form: MapSearchForm): Observable<MapApi[]> {
    return this.http.post<Map[]>(`${this.API_URL}/search`, form).pipe(
      this.waitForLookups(),
      map(maps => maps.map(map => this.populateMapFromBackend(map)))
    )
  }

  private waitForLookups<T>(): OperatorFunction<T, T> {
    return switchMap((og) => {
      return this.loading$.pipe(
        filter(loading => loading === false),
        take(1),
        map(() => og)
      );
    });
  }

  private populateMapFromBackend(map: Map): MapApi {
    const categoryLks = this.categoryLks();
    return {
      ...map,
      categories: map.categories.map(category => categoryLks?.[category])
    } as MapApi;
  }

}
