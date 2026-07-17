import { inject, Injectable, signal } from "@angular/core";
import { toSignal } from "@angular/core/rxjs-interop";
import { form, } from '@angular/forms/signals';
import { MapSearchForm } from "@arch-shared/types";

@Injectable({
  providedIn: 'root'
})
export class MapSearchFormService {

  readonly form = form(signal<MapSearchForm>({
    searchTerm: '',
    categories: [] as string[],
    onlyShowNotReviewed: false,
    onlyShowReviewed: false,
  }));

  readonly searchTerm = this.form['searchTerm']

  readonly categories = this.form['categories']



}
