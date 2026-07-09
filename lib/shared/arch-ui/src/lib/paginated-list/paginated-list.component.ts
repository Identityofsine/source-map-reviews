import { Component, computed, input, signal, TemplateRef, contentChild, OnChanges } from '@angular/core';
import { NgTemplateOutlet } from '@angular/common';
import { PaginationComponent } from '../pagination/pagination.component';
import { EmptyResultComponent } from '../empty-result/empty-result.component';

@Component({
  selector: 'arch-paginated-list',
  standalone: true,
  imports: [
    PaginationComponent,
    NgTemplateOutlet,
    EmptyResultComponent,
  ],
  templateUrl: './paginated-list.component.html',
  styleUrl: './paginated-list.component.scss'
})
export class PaginatedListComponent<T = any> implements OnChanges {
  // Inputs
  readonly data = input.required<T[]>();
  readonly itemsPerPage = input<number>(10);
  readonly showPagination = input<boolean>(true);
  readonly showPageInput = input<boolean>(true);
  readonly showPageInfo = input<boolean>(true);
  readonly layout = input<'list' | 'grid'>('list');
  readonly gridColumns = input<string>('repeat(auto-fill, minmax(280px, 1fr))');
  readonly loading = input<boolean>(false);

  // Get the template reference from content projection
  readonly itemTemplate = contentChild.required<TemplateRef<any>>('itemTemplate');

  // Internal state
  readonly currentPage = signal<number>(1);

  // Computed values
  readonly totalItems = computed(() => this.data().length);
  readonly totalPages = computed(() =>
    Math.ceil(this.totalItems() / this.itemsPerPage())
  );

  readonly paginatedData = computed(() => {
    const start = (this.currentPage() - 1) * this.itemsPerPage();
    const end = start + this.itemsPerPage();
    return this.data().slice(start, end);
  });

  onPageChange(page: number): void {
    this.currentPage.set(page);
  }

  // Reset to first page when data changes
  ngOnChanges(): void {
    this.currentPage.set(1);
  }
}
