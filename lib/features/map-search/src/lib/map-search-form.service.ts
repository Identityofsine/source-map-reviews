import { inject, Injectable, signal } from "@angular/core";
import { toSignal } from "@angular/core/rxjs-interop";
import { form, } from '@angular/forms/signals';

export type MapSearchForm = {
  searchTerm: string;
  tags: string[];
}

@Injectable({
  providedIn: 'root'
})
export class MapSearchFormService {

  readonly form = form(signal<MapSearchForm>({
    searchTerm: '',
    tags: [] as string[],
  }));

  readonly searchTerm = this.form['searchTerm']

  readonly tags = this.form['tags']



}
