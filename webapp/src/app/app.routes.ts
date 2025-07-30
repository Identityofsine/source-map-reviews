import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('@arch-feature/map-search').then(m => m.MapSearchComponent),
  }
];
