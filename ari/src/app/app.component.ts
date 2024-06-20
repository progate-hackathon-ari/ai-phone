import {Component, HostListener} from '@angular/core';
import {NavigationStart, Router} from "@angular/router";
import {GameService} from "./services/game/game.service";
import {Observable, Subject} from "rxjs";
import {NavigationService} from "./services/navigation/navigation.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent{
  title = 'ari';

  constructor(private router: Router, private navigationService: NavigationService) {
    this.navigationService.previousUrl = document.referrer;

    this.router.events.subscribe(event => {
      if (event instanceof NavigationStart) {
        if (event.navigationTrigger === 'popstate') {
          // 戻るボタンが押されたときの処理
          const previousUrl = this.navigationService.getPreviousUrl();
          if (previousUrl) {
            this.router.navigateByUrl(previousUrl).then()
          }
        }
      }
    });
  }
}
