import { inject, Injectable } from "@angular/core";
import { toSignal } from "@angular/core/rxjs-interop";
import { FormBuilder } from "@angular/forms";

@Injectable({
  providedIn: 'root'
})
export class MapSearchFormService {

  readonly fb = inject(FormBuilder);

  readonly form = this.fb.group({
    searchTerm: [''],
    tags: [[] as string[]]
  });

  readonly searchTerm = toSignal(
    this.form?.get('searchTerm')!.valueChanges
  );

  readonly tags = toSignal(
    this.form?.get('tags')!.valueChanges
  );



}
