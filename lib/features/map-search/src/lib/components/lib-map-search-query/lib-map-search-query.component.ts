import { Component, computed, effect, inject, resource, untracked } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MapSearchFormService } from '../../map-search-form.service';
import { LookupCacheService, MapsService } from '@arch-shared/data-source';
import { InputTextModule } from 'primeng/inputtext';
import { FormField } from "@angular/forms/signals";
import { SelectModule } from "primeng/select";

@Component({
  selector: 'arch-map-search-query',
  templateUrl: './lib-map-search-query.component.html',
  styleUrls: ['./lib-map-search-query.component.scss'],
  imports: [
    ReactiveFormsModule,
    FormsModule,
    InputTextModule,
    FormField,
    SelectModule,
  ],
  standalone: true,
})
export class LibMapSearchQueryComponent {
  readonly mapsService = inject(MapsService);
  readonly formService = inject(MapSearchFormService);
  readonly form = this.formService.form;
  readonly lookupsService = inject(LookupCacheService);

  readonly categoryLks = this.lookupsService.getLookup('mapCategoryLk');

  readonly tags = computed(() => {
    const categoryLks = this.categoryLks();
    if (!categoryLks) {
      return [];
    }
    return categoryLks.map(tag => ({
      label: tag.shortDescription,
      value: tag.lkMapCategory
    }));
  });

}
