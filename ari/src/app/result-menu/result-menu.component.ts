import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService, dataSubscribe} from "../services/game/game.service";
import {Observable, Subscription} from "rxjs";

interface Player {
    id: number;
    name: string;
    prompt: string;
    image_uri: string;
  }

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
  //仮おき
    players: Player[] = [
      { id: 1, name: 'Player 1', prompt: '卵を焼いているフランス人',image_uri:"assets/sample.jpeg"},
      { id: 2, name: 'Player 2', prompt: 'world' ,image_uri:"assets/sample.jpeg"},
      { id: 3, name: 'Player 3', prompt: 'hogehoge' ,image_uri:"assets/sample.jpeg"},
    ];
    selectedPlayer: Player = this.players[0];
  selectPlayer(player: Player): void {
    this.selectedPlayer = player;
  }
}
