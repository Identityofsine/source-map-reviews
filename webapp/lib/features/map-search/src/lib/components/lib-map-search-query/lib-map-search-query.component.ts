import { Component, inject } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { ArchTextInputComponent, ArchDropdownInputComponent, DropdownItem } from '@arch-shared/arch-ui';
import { MapSearchFormService } from '../../map-search-form.service';

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
            formControlName="tags"
            label="Tags"
            placeholder="Select or add tags..."
            [freeRange]="true"
            [items]="tagItems"
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
  readonly formService = inject(MapSearchFormService);
  readonly form = this.formService.form;
  
  // Common tags that users can select from
  readonly tagItems: DropdownItem[] = [
    { key: 'action', value: 'Action' },
    { key: 'adventure', value: 'Adventure' },
    { key: 'puzzle', value: 'Puzzle' },
    { key: 'horror', value: 'Horror' },
    { key: 'comedy', value: 'Comedy' },
    { key: 'multiplayer', value: 'Multiplayer' },
    { key: 'singleplayer', value: 'Single Player' },
    { key: 'campaign', value: 'Campaign' },
    { key: 'survival', value: 'Survival' },
    { key: 'rpg', value: 'RPG' },
    { key: 'strategy', value: 'Strategy' },
    { key: 'tower-defense', value: 'Tower Defense' },
    { key: 'racing', value: 'Racing' },
    { key: 'platformer', value: 'Platformer' },
    { key: 'shooter', value: 'Shooter' },
  ];
} 