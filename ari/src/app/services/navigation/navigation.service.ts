import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class NavigationService {
  previousUrl: string;

  constructor() {
    this.previousUrl = document.referrer;
  }

  public getPreviousUrl(): string {
    return this.previousUrl;
  }
}
