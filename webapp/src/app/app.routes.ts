import { Routes } from '@angular/router';
import { LOGIN_PATH } from '@arch-shared/types';

export const routes: Routes = [
  {
    path: 'search',
    loadComponent: () => import('@arch-feature/map-search').then(m => m.MapSearchComponent),
  },
  {
    path: 'map/:id',
    loadComponent: () => import('@arch-feature/maps').then(m => m.MapsComponent),
  },
  {
    path: LOGIN_PATH,
    loadComponent: () => import('@arch-feature/auth').then(m => m.LoginComponent),
  }
];
