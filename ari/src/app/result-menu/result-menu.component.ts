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
        username:string;
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
interface DataList{
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
                        "img": "https://d14ubbdtfe7bkh.cloudfront.net/0190407b-703b-2b490b/8fc3464e-5d37-4ac4-bfea-5334251f81e4/0.jpg",
                        "username": "user1"
                    },
                    "1": {
                        "answer": "what",
                        "img": "https://d14ubbdtfe7bkh.cloudfront.net/0190407b-703b-2b490b/8fc3464e-5d37-4ac4-bfea-5334251f81e4/1.jpg",
                        "username": "user1"
                    }
                }
            },
            "bbbbbbbbb": {
                "img_score": 64,
                "per_user": {
                    "0": {
                        "answer": "d",
                        "img": "https://d14ubbdtfe7bkh.cloudfront.net/0190407b-703b-2b490b/9a41f207-6d79-4c4d-bf4c-e3254f2029f2/0.jpg",
                        "username": "user1"
                    },
                    "1": {
                        "answer": "car",
                        "img": "https://d14ubbdtfe7bkh.cloudfront.net/0190407b-703b-2b490b/9a41f207-6d79-4c4d-bf4c-e3254f2029f2/1.jpg",
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
                        console.log("img_score 2 " + score);
                        console.log("key  " + userKey);
                        console.log("value  " + userValue.answer);
                        console.log("image  " + userValue.img);
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
    
            //   console.log(this.players)
            })
            
                this.selectedPlayer = this.players[0]
              }
    }
)

    //   let json = JSON.parse(data)

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
}

