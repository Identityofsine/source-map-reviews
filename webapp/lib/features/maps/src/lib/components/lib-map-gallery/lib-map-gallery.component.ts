import { ChangeDetectionStrategy, Component, computed, effect, inject, input, signal, ViewContainerRef } from '@angular/core';
import { ArchContainer } from '@arch-shared/arch-ui';
import { MapImage, MapReview } from '@arch-shared/types';
import { MapGalleryReviewComponent } from './lib-map-gallery-review/lib-map-gallery-review.component';
import { AuthService } from '@arch-shared/auth';
import { BehaviorSubject, catchError, of } from 'rxjs';
import { toSignal } from '@angular/core/rxjs-interop';
import { HttpClient } from '@angular/common/http';
import { MapImageDetailsComponent } from '../lib-map-image-details/lib-map-image-details.component';

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
  readonly containerRef = inject(ViewContainerRef);

  readonly selectedReview = input<MapReview | null>(null);

  readonly mapName = input<string>();
  readonly mapImages = input<MapReview[]>();

  readonly canAddImage = computed(() => {
    return this.externAuthService.isAuthenticatedSignal();
  });

  readonly imagesFromMapImages = computed(() => {
    return this.mapImages()?.map((review) => review.images).flat() ?? [];
  });

  readonly allImages$ = new BehaviorSubject<MapImage[]>(this.imagesFromMapImages());

  readonly allImages = toSignal(this.allImages$.asObservable())

  //fill in the gaps of images with a review, should equal the size of allImages but be filled in with nulls where there are no reviews
  readonly reviewImages = computed(() => {
    const allImages = this.allImages() ?? [];
    const mapImages = this.mapImages() ?? [];
    const reviewImages: (MapReview | null)[] = [];

    for (let i = 0; i < allImages.length; i++) {
      const image = allImages[i];
      const review = mapImages.find(review => review.images.some(img => img.imageId === image.imageId));
      reviewImages.push(review ?? null);
    }

    return reviewImages;
  });

  readonly isEmpty = computed(() => {
    return (this.allImages() ?? []).length <= 0;
  });

  readonly currentReview = computed(() => {
    return this.reviewImages()?.[this.currentIndex()] ?? null;
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
      const curReview = this.selectedReview();
      if (!curReview) {
        this.currentIndex.set(0);
        return;
      }
      const idx = this.allImages()?.findIndex(i => i.mapReviewId === curReview.mapReviewId) ?? 0;
      this.currentIndex.set(idx);
    })

    effect(() => {
      const mapImages = this.imagesFromMapImages();
      this.allImages$.next(mapImages);
    })

    effect(() => {
      const mapName = this.mapName();
      const mapImages = this.imagesFromMapImages();
      this.http.get<any>(`/api/images/${this.mapName()}.jpg`)
        .pipe(
          catchError((e) => of(e)),
        )
        .subscribe((blob) => {
          if (blob.status === 200) {
            const newArray = [
              {
                image: {
                  imageId: -1,
                  imagePath: `/api/images/${mapName}.jpg`,
                  caption: 'Fetched from gamebanana',
                },
              },
              ...mapImages
            ]
            this.allImages$.next(newArray);
            console.log('Fetched images from gamebanana', newArray);
          }
        });
    })
  }

  public openAddImage() {
    const view = this.containerRef.createComponent(MapImageDetailsComponent);
    //view.setInput();
    view.instance.shouldReloadImage.subscribe((res) => {
      if (res) {
        this.reloadImages();
      }
      view.destroy();
    });
  }

  public reloadImages() {
    const mapImages = this.imagesFromMapImages();
    this.allImages$.next(mapImages);
  }


}
