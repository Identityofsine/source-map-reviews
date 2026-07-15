import { Component, computed, effect, inject, untracked } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ArchTextInputComponent, ArchDropdownInputComponent, DropdownItem } from '@arch-shared/arch-ui';
import { MapSearchFormService } from '../../map-search-form.service';
import { rxResource } from '@angular/core/rxjs-interop';
import { MapsService } from '@arch-shared/data-source';
import { InputTextModule } from 'primeng/inputtext';
import { FormField } from "@angular/forms/signals";
import { SelectModule } from "primeng/select";
import { TagLkApi } from '@arch-shared/types';

@Component({
  selector: 'arch-map-search-query',
  template: `
    <div class="map-search-query" [formGroup]="form">
      <form>
        <div class="search-controls items-center">
          <div class="search-input-group">
            <input type="text" pInputText [formField]="form.searchTerm" placeholder="Search maps..." />
          </div>
          <div class="tags-input-group w-full h-[53px] ">
            <p-select 
              [options]="tags()"
              class="w-full h-full items-center"
              [formField]="form.tags"
              optionLabel="tagDescription"
              optionValue="tagLk"
              placeholder="Select tags"
            />
          </div>
        </div>
      </form>
    </div>
  `,
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

  readonly tagsLks = rxResource({
    stream: () => this.mapsService.getTags(),
  });

  readonly tags = computed(() => [{
    tagLk: 'defuse',
    tagDescription: 'Defuse',
  }] as TagLkApi[]);

  constructor() {
    effect(() => {
      const tags = this.tags();
      untracked(() => {
        this.form.tags().value.set(tags.map(tag => tag.tagLk));
      })
    })
  }
}
