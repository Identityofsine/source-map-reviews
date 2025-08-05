import { ChangeDetectionStrategy, Component, computed, inject, input, signal } from '@angular/core';
import { ArchContainer } from '@arch-shared/arch-ui';
import { MapReview } from '@arch-shared/types';
import { MapGalleryReviewComponent } from './lib-map-gallery-review/lib-map-gallery-review.component';
import { AuthService } from '@arch-shared/auth';

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

  readonly externAuthService = inject(AuthService);
  readonly mapImages = input<MapReview[]>();

  readonly isEmpty = computed(() => {
    return (this.mapImages() ?? []).length <= 0;
  });

  readonly canAddImage = computed(() => {
    return this.externAuthService.isAuthenticatedSignal();
  });

  readonly allImages = computed(() => {
    return this.mapImages()
      ?.map((review) => review.images)
      ?.flatMap(i => i)
  })

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



}
