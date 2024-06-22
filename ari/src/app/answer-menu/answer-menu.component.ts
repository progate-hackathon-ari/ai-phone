import {Component, OnDestroy, OnInit} from '@angular/core';
import {GameService, dataSubscribe} from "../services/game/game.service";
import {Router} from "@angular/router";
import {Observable, Subscription} from "rxjs";

interface AnswerData {
  image_uri: string;
}

@Component({
  selector: 'app-answer-menu',
  templateUrl: './answer-menu.component.html',
  styleUrl: './answer-menu.component.scss'
})
export class AnswerMenuComponent implements OnInit, OnDestroy{
  constructor(private router: Router, private gameService: GameService, private dataSubscribe: dataSubscribe) {
  }
  dsub: Observable<any> | undefined;
  Subs: Subscription | undefined;
  imageUri: string = '';
  isButtonVisible = true;
  isButtonEnabled = true;
  countDown = "30";

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    this.gameService.sendCountDown(30)

    this.dsub = this.dataSubscribe.subscribe();

    this.Subs = this.dsub.subscribe(data => {
      const json = JSON.parse(data)

      if (json.is_done != undefined) {
        this.countDown = ( '000' + json.count ).slice( -2 );
        if (json.is_done && this.isButtonVisible){
          this.onClickSubmit()
        }
      }

      if (json.is_all_user_answered) {
        if (this.gameService.isAdmin){
          this.gameService.sendNext()
        }
      }else if (json.state === "next_round") {
        if (!json.data){
          this.router.navigateByUrl('/answer').then();
          return;
        }
      }else if(json.state === "game_end") {
        this.router.navigateByUrl('/result').then()
      }
    })

  }

  ngOnDestroy(): void {
    if (this.Subs) {
      this.Subs.unsubscribe();
    }
  }

  question: string = ''; // デフォルト値を設定

  onChangeQuestion(event:any): void {
    this.question = event.target.value;
  }

  onClickSubmit(){
    console.log(this.question)
    this.gameService.sendAnswer(this.question)
    this.isButtonVisible = false;
    this.isButtonEnabled = false;
  }
}
