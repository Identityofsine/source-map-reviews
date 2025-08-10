import { ChangeDetectionStrategy, Component, computed, effect, inject, input, signal } from '@angular/core';
import { ArchContainer } from '@arch-shared/arch-ui';
import { MapImage, MapReview } from '@arch-shared/types';
import { MapGalleryReviewComponent } from './lib-map-gallery-review/lib-map-gallery-review.component';
import { AuthService } from '@arch-shared/auth';
import { BehaviorSubject, catchError, of } from 'rxjs';
import { takeUntilDestroyed, toSignal } from '@angular/core/rxjs-interop';
import { HttpClient } from '@angular/common/http';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-map-gallery',
  templateUrl: './lib-map-gallery.component.html',
  styleUrl: './lib-map-gallery.component.scss',
  imports: [
    ArchContainer,
    MapGalleryReviewComponent
  ],
})
export class MapGalleryComponent {

  readonly http = inject(HttpClient);
  readonly externAuthService = inject(AuthService);
  readonly mapName = input<string>();
  readonly mapImages = input<MapReview[]>();

  readonly canAddImage = computed(() => {
    return this.externAuthService.isAuthenticatedSignal();
  });

  readonly allImages$ = new BehaviorSubject<MapImage[]>([]);

  readonly allImages = toSignal(this.allImages$.asObservable())


  readonly isEmpty = computed(() => {
    return (this.allImages() ?? []).length <= 0;
  });

  readonly currentReview = computed(() => {
    return this.mapImages()?.[this.currentIndex()] ?? null;
  })

  readonly currentImage = computed(() => {
    return this.allImages()?.[this.currentIndex()] ?? null;
  })

  readonly currentImagePath = computed(() => {
    return this.currentImage()?.image?.imagePath;
  });

  readonly currentIndex = signal(0);

  constructor() {
    effect(() => {
      const mapName = this.mapName();
      this.http.get<any>(`/api/images/${this.mapName()}.jpg`)
        .pipe(
          catchError((e) => of(e)),
        )
        .subscribe((blob) => {
          if (blob.status === 200) {
            this.allImages$.next([
              {
                image: {
                  imageId: -1,
                  imagePath: `/api/images/${mapName}.jpg`,
                  caption: 'Fetched from gamebanana',
                },
              }
            ])
          }
        });
    })
  }


}
