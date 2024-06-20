import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class NavigationService {
  previousUrl: string | undefined;

  public getPreviousUrl(): string {
    if (this.previousUrl === undefined) {
      return '/';
    }
    return this.previousUrl;
  }
}
