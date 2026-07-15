import { Component, computed, DestroyRef, inject, Injector, input, OnInit, Signal } from '@angular/core';
import { rxResource, takeUntilDestroyed, toObservable, toSignal } from '@angular/core/rxjs-interop';
import { ReactiveFormsModule } from '@angular/forms';
import { MapsService } from '@arch-shared/data-source';
import { PaginatedListComponent } from '@arch-shared/arch-ui';
import { MapThumbnailComponent } from './components/lib-map-thumbnail/lib-map-thumbnail.component';
import { LibMapSearchQueryComponent } from './components/lib-map-search-query/lib-map-search-query.component';
import { ActivatedRoute, Router } from '@angular/router';
import { combineLatest, map, of, startWith } from 'rxjs';
import { MapSearchForm, MapSearchFormService } from './map-search-form.service';
import { MessageModule } from 'primeng/message';

@Component({
  selector: 'arch-map-search',
  imports: [
    LibMapSearchQueryComponent,
    ReactiveFormsModule,
    MapThumbnailComponent,
    PaginatedListComponent,
    MessageModule,
  ],
  templateUrl: './map-search.component.html',
  styleUrl: './map-search.component.scss',
})
export class MapSearchComponent implements OnInit {

  readonly route = inject(ActivatedRoute);
  readonly injector = inject(Injector);
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

  readonly searchFormBuilt = computed(() => {
    return this.form().value();
  })

  readonly isDirty = computed(() => this.form().dirty);

  readonly search = rxResource({
    params: () => ({
      searchForm: this.searchFormBuilt(),
    }),
    stream: ({ params }) => {
      if (!params.searchForm) {
        return of(undefined);
      }
      return this.mapService.searchMaps({
        ...params.searchForm,
      })
    }
  });

  readonly maps = computed(() => {
    const data = this.search.value() ?? [];
    return data; // Return all data, let pagination component handle slicing
  })

  ngOnInit(): void {
    this.queryParamsToForm();

    toObservable(computed(() => this.form().value()), { injector: this.injector })
      .pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(({ searchTerm, tags }) => {
        this.router.navigate([], {
          relativeTo: this.route,
          queryParams: {
            _search: searchTerm || null,
            _tags: Array.isArray(tags) && tags.length > 0
              ? tags.join(this.SEPARATOR)
              : null,
          },
          queryParamsHandling: 'merge',
        });
      });
  }
  private queryParamsToForm() {
    if (this._search()) {
      this.form.searchTerm().value.set(this._search());
    }
    if (this._tags()) {
      this.form.tags().value.set(this._tags().split(this.SEPARATOR));
    }
  }

}
