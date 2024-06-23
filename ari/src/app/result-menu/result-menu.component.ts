import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService, dataSubscribe} from "../services/game/game.service";
import {Observable, Subscription} from "rxjs";

interface Player {
    key: string;
    id: string;
    score: number | undefined;
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
    key: string;
  prompt: string;
  image_uri: string;
  user: string;
  score?: number;
}


const data = `{
    "state": "game_end",
    "data": {
        "result": {
            "aaa": {
                "img_score": 52,
                "per_user": {
                    "0": {
                        "answer": "No answer",
                        "img": "assets/sample.jpeg",
                        "username": "user1"
                    },
                    "1": {
                        "answer": "what",
                        "img": "assets/sample.jpeg",
                        "username": "user1"
                    }
                }
            },
            "bbbbbbbbb": {
                "img_score": 64,
                "per_user": {
                    "0": {
                        "answer": "d",
                        "img": "assets/sample.jpeg",
                        "username": "user1"
                    },
                    "1": {
                        "answer": "car",
                        "img": "assets/sample.jpeg",
                        "username": "user1"
                    }
                }
            }
        }
    }
}`;

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
        
                console.log(value);
            if (value != undefined) {
                if ('img_score' in value) {
                    score = value.img_score as number;
                    console.log("img_score 1 " + score);
                }
                if ('per_user' in value) {
                    Object.entries(value.per_user as PerUser).forEach(([userKey, userValue]) => {
                        console.log(userKey);
                        listData.push({
                            key: userKey,
                            prompt: userValue.answer,
                            image_uri: userValue.img,
                            user: userValue.username,
                            score: score 
                        });
                    });
                }
                    this.players.push({
                        key: key,
                      id: key,
                      score: score,
                      data: listData
                    })
            }
            //   console.log(this.players)
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

