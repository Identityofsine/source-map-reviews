import { HttpClient } from "@angular/common/http"
import { computed, inject, Injectable, Service, Signal } from "@angular/core"
import { rxResource } from "@angular/core/rxjs-interop";
import { BaseLookups, LookupKeys, LookupMap } from "@arch-shared/types";
import { Observable } from "rxjs"
import { composeLookups } from "./util/lookups.util";

@Injectable({
  providedIn: 'root',
})
export class LookupsService {
  readonly API_URL = `/api/lookups`
  readonly http = inject(HttpClient);

  public getLookups(): Observable<LookupMap> {
    return this.http.get<LookupMap>(`${this.API_URL}/all`);
  }
}

@Service()
export class LookupCacheService {

  readonly lookupsHttpService = inject(LookupsService)

  readonly lookupsRequest = rxResource({
    stream: () => this.lookupsHttpService.getLookups(),
  });

  private readonly _lookups = computed(() => this.lookupsRequest.value());

  public loading = this.lookupsRequest.isLoading;

  private readonly _composedLookups = computed(() => {
    const lookupsRaw = this._lookups();
    if (!lookupsRaw) {
      return null;
    }
    return composeLookups(lookupsRaw);
  });

  public getLookup<T extends LookupKeys>(key: T): Signal<BaseLookups[T] | null> {
    return computed(() => {
      const lookups = this._composedLookups();
      if (!lookups) {
        return null;
      }
      return lookups?.lookups?.[key];
    })
  }

}

