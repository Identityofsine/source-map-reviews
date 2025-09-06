import { inject, Pipe, PipeTransform, OnDestroy } from '@angular/core';
import { UserService } from '@arch-shared/data-source';
import { Subscription } from 'rxjs';

@Pipe({
  name: 'username',
  pure: false
})
export class UsernamePipe implements PipeTransform, OnDestroy {

  private readonly userService = inject(UserService);
  private readonly usernameCache = new Map<number, string>();
  private readonly subscriptions = new Map<number, Subscription>();

  transform(value: number): string {
    if (!value) {
      return 'Unknown';
    }

    // Return cached value if available
    if (this.usernameCache.has(value)) {
      return this.usernameCache.get(value)!;
    }

    // Don't fetch if already fetching
    if (this.subscriptions.has(value)) {
      return 'Loading...';
    }

    // Set loading state and fetch username
    this.usernameCache.set(value, 'Loading...');
    
    const subscription = this.userService.getUsername(value).subscribe({
      next: (username) => {
        this.usernameCache.set(value, username || 'Unknown');
        this.subscriptions.delete(value);
      },
      error: () => {
        this.usernameCache.set(value, 'Unknown');
        this.subscriptions.delete(value);
      }
    });
    
    this.subscriptions.set(value, subscription);
    return 'Loading...';
  }

  ngOnDestroy(): void {
    // Clean up all subscriptions
    this.subscriptions.forEach(sub => sub.unsubscribe());
    this.subscriptions.clear();
    this.usernameCache.clear();
  }
}
