import { Component, computed, input, output, signal, OnChanges, DestroyRef, inject } from '@angular/core';
import { FormsModule, FormControl, ReactiveFormsModule } from '@angular/forms';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { ArchDropdownInputComponent } from '../dropdown-input/dropdown-input.component';

@Component({
  selector: 'arch-pagination',
  standalone: true,
  imports: [FormsModule, ReactiveFormsModule, ArchDropdownInputComponent],
  templateUrl: './pagination.component.html',
  styleUrl: './pagination.component.scss'
})
export class PaginationComponent implements OnChanges {
  private readonly destroyRef = inject(DestroyRef);
  
  // Inputs
  readonly totalItems = input.required<number>();
  readonly itemsPerPage = input<number>(25);
  readonly currentPage = input<number>(1);
  readonly showPageInput = input<boolean>(true);
  readonly showPageInfo = input<boolean>(true);
  readonly showItemsPerPageDropdown = input<boolean>(true);
  readonly itemsPerPageOptions = input<number[]>([25, 50, 100, 250]);

  // Outputs
  readonly pageChange = output<number>();
  readonly itemsPerPageChange = output<number>();

  // Internal state for page input
  readonly pageInputValue = signal<number>(1);
  
  // Form control for items per page dropdown
  readonly itemsPerPageControl = new FormControl<string[]>([]);
  
  // Computed dropdown items
  readonly itemsPerPageDropdownItems = computed(() => 
    this.itemsPerPageOptions().map(value => ({
      key: value.toString(),
      value: `${value} per page`
    }))
  );

  // Computed values
  readonly totalPages = computed(() => 
    Math.ceil(this.totalItems() / this.itemsPerPage())
  );

  readonly canGoPrevious = computed(() => this.currentPage() > 1);
  readonly canGoNext = computed(() => this.currentPage() < this.totalPages());

  readonly startItem = computed(() => 
    (this.currentPage() - 1) * this.itemsPerPage() + 1
  );

  readonly endItem = computed(() => 
    Math.min(this.currentPage() * this.itemsPerPage(), this.totalItems())
  );

  constructor() {
    // Sync page input with current page
    this.pageInputValue.set(this.currentPage());
    // Set initial items per page value
    this.itemsPerPageControl.setValue([this.itemsPerPage().toString()]);
    
    // Subscribe to items per page changes
    this.itemsPerPageControl.valueChanges
      .pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(value => this.onItemsPerPageChange(value || []));
  }

  ngOnChanges() {
    this.pageInputValue.set(this.currentPage());
    // Update dropdown when itemsPerPage changes
    this.itemsPerPageControl.setValue([this.itemsPerPage().toString()]);
  }

  goToPage(page: number): void {
    if (page >= 1 && page <= this.totalPages()) {
      this.pageChange.emit(page);
      this.pageInputValue.set(page);
    }
  }

  goToPrevious(): void {
    if (this.canGoPrevious()) {
      this.goToPage(this.currentPage() - 1);
    }
  }

  goToNext(): void {
    if (this.canGoNext()) {
      this.goToPage(this.currentPage() + 1);
    }
  }

  goToFirst(): void {
    this.goToPage(1);
  }

  goToLast(): void {
    this.goToPage(this.totalPages());
  }

  onPageInputChange(value: string): void {
    const pageNumber = parseInt(value, 10);
    if (!isNaN(pageNumber)) {
      this.pageInputValue.set(pageNumber);
    }
  }

  onPageInputEnter(): void {
    this.goToPage(this.pageInputValue());
  }

  onPageInputBlur(): void {
    // Reset to current page if invalid input
    if (this.pageInputValue() < 1 || this.pageInputValue() > this.totalPages()) {
      this.pageInputValue.set(this.currentPage());
    }
  }

  onItemsPerPageChange(selectedValues: string[]): void {
    if (selectedValues && selectedValues.length > 0) {
      const itemsPerPage = parseInt(selectedValues[0], 10);
      if (!isNaN(itemsPerPage) && this.itemsPerPageOptions().includes(itemsPerPage)) {
        this.itemsPerPageChange.emit(itemsPerPage);
        // Reset to first page when changing items per page
        this.goToPage(1);
      }
    }
  }
} 