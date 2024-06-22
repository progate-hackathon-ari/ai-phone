import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService, dataSubscribe} from "../services/game/game.service";
import {Observable, Subscription} from "rxjs";

interface Player {
  id: string;
  data: DataList[];
}

interface PerUser {
  [key: string]: {
    img: string;
    answer: string;
    username: string;
  };
}

interface ImgScore {
  per_user: PerUser;
  score: number;
}

interface Result {
  [key: string]: ImgScore;
}

interface DataStructure {
  state: string;
  data: {
    result: Result;
  };
}

interface DataList {
  prompt: string;
  image_uri: string;
  user: string;
  score?: number;
}

@Component({
  selector: 'app-result-menu',
  templateUrl: './result-menu.component.html',
  styleUrl: './result-menu.component.scss'
})
export class ResultMenuComponent implements OnInit, OnDestroy {
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
        const json: DataStructure = JSON.parse(data);
        console.log(json);
        if (json.state === "game_end" && json.data) {
          Object.entries(json.data.result).forEach(([key, value]) => {
            let listData: DataList[] = [];
            let score: number | undefined;

            if (value != undefined) {
              if ('img_score' in value) {
                score = value.img_score as number;
              }
              if ('per_user' in value) {
                Object.entries(value.per_user as PerUser).forEach(([userKey, userValue]) => {
                  listData.push({
                    prompt: userValue.answer,
                    image_uri: userValue.img,
                    user: userValue.username,
                    score: score // img_score を同じユーザーのデータに含める
                  });
                });
              }
              this.players.push({
                id: key,
                data: listData
              })
            }

          })

          this.selectedPlayer = this.players[0]
        }
      }
    )

  }

  ngOnDestroy(): void {
    if (this.Subs) {
      this.Subs.unsubscribe();
    }
  }

  selectPlayer(player: Player): void {
    this.selectedPlayer = player;
    console.log(this.selectedPlayer);
  }

  onClickHome(){
    this.router.navigateByUrl(`${document.location.origin}/`).then()
  }
}

