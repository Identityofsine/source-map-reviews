import { Component, computed, DestroyRef, inject, input, OnInit } from '@angular/core';
import { rxResource, takeUntilDestroyed, toSignal } from '@angular/core/rxjs-interop';
import { ReactiveFormsModule } from '@angular/forms';
import { MapsService } from '@arch-shared/data-source';
import { PaginatedListComponent } from '@arch-shared/arch-ui';
import { MapThumbnailComponent } from './components/lib-map-thumbnail/lib-map-thumbnail.component';
import { LibMapSearchQueryComponent } from './components/lib-map-search-query/lib-map-search-query.component';
import { ActivatedRoute, Router } from '@angular/router';
import { combineLatest, map, startWith } from 'rxjs';
import { MapSearchFormService } from './map-search-form.service';

@Component({
  selector: 'arch-map-search',
  imports: [
    LibMapSearchQueryComponent,
    ReactiveFormsModule,
    MapThumbnailComponent,
    PaginatedListComponent,
  ],
  templateUrl: './map-search.component.html',
  styleUrl: './map-search.component.scss',
})
export class MapSearchComponent implements OnInit {

  readonly route = inject(ActivatedRoute);
  readonly formService = inject(MapSearchFormService);
  readonly router = inject(Router);
  readonly mapService = inject(MapsService);
  readonly destroyRef = inject(DestroyRef);

  readonly SEPARATOR = '>';

  //?_search=""
  readonly _search = input<string>();
  //?_tags=""
  readonly _tags = input<string>();

  readonly form = this.formService.form;
  readonly searchTerm = this.formService.searchTerm;
  readonly tags = this.formService.tags;

  readonly search = rxResource({
    request: () => ({
      searchTerm: this.searchTerm(),
      tags: this.tags() ?? [],
    }),
    loader: () => this.mapService.searchMaps({
      searchTerm: this.searchTerm() ?? '',
      tags: [this.tags() ?? []].flat(),
    }),
  });

  readonly maps = computed(() => {
    const data = this.search.value() ?? [];
    return data; // Return all data, let pagination component handle slicing
  })

  ngOnInit(): void {

    this.queryParamsToForm();

    const values = Object.keys(this.form?.value)
      ?.map(key =>
        this.form.get(key)?.valueChanges
          .pipe(
            startWith(this.form.get(key)?.value),
            map((value) => ({ [key]: value }))
          )
      );

    combineLatest(values)
      .pipe(
        takeUntilDestroyed(this.destroyRef),
        map((values) => {
          return values.reduce((acc, curr) => {
            return { ...acc, ...curr };
          }, {});
        })
      )
      .subscribe(({ searchTerm, tags }) => {
        this.router.navigate([], {
          relativeTo: this.route,
          queryParams: {
            _search: searchTerm,
            _tags: tags.length > 0 &&
              Array.isArray(tags) ?
              tags.join(this.SEPARATOR)
              : null,
          },
          queryParamsHandling: 'merge',
        });
      });
  }

  private queryParamsToForm() {
    if (this._search()) {
      this.form.get('searchTerm')?.setValue(this._search(), { emitEvent: true });
    }
    if (this._tags()) {
      this.form.get('tags')?.setValue(this._tags().split(this.SEPARATOR), { emitEvent: true });
    }
  }

}
