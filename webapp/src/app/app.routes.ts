import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: 'search',
    loadComponent: () => import('@arch-feature/map-search').then(m => m.MapSearchComponent),
  },
  {
    path: 'map/:id',
    loadComponent: () => import('@arch-feature/maps').then(m => m.MapsComponent),
  }
];
