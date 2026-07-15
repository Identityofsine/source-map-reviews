import { HttpClient } from "@angular/common/http"
import { inject, Injectable } from "@angular/core"
import { Lookups } from "@arch-shared/types"
import { LookupMap } from "lib/shared/types/src/lib/lookup.types"
import { Observable } from "rxjs"

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
