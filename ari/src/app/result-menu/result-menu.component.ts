import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService, dataSubscribe} from "../services/game/game.service";
import {Observable, Subscription} from "rxjs";

interface Player {
    id: string;
    data: DataList[];
}


interface DataList{
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
  players: Player[] = []
  selectedPlayer: Player = this.players[0];

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    this.dsub = this.dataSubscribe.subscribe();

    this.Subs = this.dsub.subscribe(data => {
      let json = JSON.parse(data)

      if(json.state === "game_end" && json.data) {

        Object.entries(json.data.result).forEach(([key, value]) => {
          let listData: DataList[] = []

          if (value != undefined) {
            Object.entries(value).forEach(([key2, value2]) => {
              listData.push({
                prompt: value2.prompt,
                image_uri: value2.image_uri
              })
            })

            this.players.push({
              id: key,
              data: listData
            })
          }
          console.log(this.players)
        })
        this.selectedPlayer = this.players[0]
      }
    })
  }

  ngOnDestroy(): void {
    if (this.Subs) {
      this.Subs.unsubscribe();
    }
  }

  selectPlayer(player: Player): void {
    this.selectedPlayer = player;
  }
}

