import { Component, computed, inject } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { ArchTextInputComponent, ArchDropdownInputComponent, DropdownItem } from '@arch-shared/arch-ui';
import { MapSearchFormService } from '../../map-search-form.service';
import { rxResource } from '@angular/core/rxjs-interop';
import { MapsService } from '@arch-shared/data-source';

@Component({
  selector: 'arch-map-search-query',
  template: `
    <div class="map-search-query" [formGroup]="form">
      <div class="search-controls">
        <div class="search-input-group">
          <arch-text-input
            formControlName="searchTerm"
            label="Search Maps"
            placeholder="Enter map name or description..."
          />
        </div>

        <div class="tags-input-group">
          <arch-dropdown-input
            [items]="tagsLks.value()"
            itemKey="tagLk"
            itemValue="tagLk"
            formControlName="tags"
            label="Tags"
            placeholder="Select or add tags..."
            [freeRange]="true"
          />
        </div>
      </div>
    </div>
  `,
  styleUrls: ['./lib-map-search-query.component.scss'],
  imports: [
    ReactiveFormsModule,
    ArchTextInputComponent,
    ArchDropdownInputComponent,
  ],
  standalone: true,
})
export class LibMapSearchQueryComponent {
  readonly mapsService = inject(MapsService);
  readonly formService = inject(MapSearchFormService);
  readonly form = this.formService.form;

  readonly tagsLks = rxResource({
    loader: () => this.mapsService.getTags(),
  });

  readonly tags = computed(() => this.tagsLks.value() ?? []);
}
