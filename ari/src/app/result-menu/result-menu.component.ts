import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService, dataSubscribe} from "../services/game/game.service";
import {Observable, Subscription} from "rxjs";

@Component({
  selector: 'app-result-menu',
  templateUrl: './result-menu.component.html',
  styleUrl: './result-menu.component.scss'
})
export class ResultMenuComponent implements OnInit, OnDestroy{
  constructor(private router: Router, private gameService: GameService, private dataSubscribe: dataSubscribe) {
  }
  dsub: Observable<any> | undefined;
  Subs: Subscription | undefined;

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    this.dsub = this.dataSubscribe.subscribe();

    this.Subs = this.dsub.subscribe(data => {
      let json = JSON.parse(data)
      console.log(json)
    })
  }

  ngOnDestroy(): void {
    if (this.Subs) {
      this.Subs.unsubscribe();
    }
  }
}
