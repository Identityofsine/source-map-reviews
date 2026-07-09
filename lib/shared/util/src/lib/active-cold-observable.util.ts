import { Observable, ObservableInput } from "rxjs";

export function awakeColdObservable<T>(cold: () => Observable<T>): () => ObservableInput<T> {
  return () => new Observable<T>(subscriber => {
    const subscription = cold().subscribe(subscriber);
    return () => {
      subscription.unsubscribe();
    };
  });
}
