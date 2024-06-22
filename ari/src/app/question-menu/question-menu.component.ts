import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService,dataSubscribe} from "../services/game/game.service";
import {Observable, Subscription} from "rxjs";

@Component({
  selector: 'app-question-menu',
  templateUrl: './question-menu.component.html',
  styleUrl: './question-menu.component.scss'
})
export class QuestionMenuComponent implements OnInit, OnDestroy{
  constructor(private router: Router, private gameService: GameService, private dataSubscribe: dataSubscribe) {
  }
  dsub: Observable<any> | undefined;
  Subs: Subscription | undefined;
  isButtonEnabled = true;
  countDown = "30";

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    this.gameService.sendCountDown(10)

    this.dsub = this.dataSubscribe.subscribe();

    this.Subs = this.dsub.subscribe(data => {
        const json = JSON.parse(data)

        if (json.is_done != undefined) {
          this.countDown = ( '000' + json.count ).slice( -2 );
          if (json.is_done && this.isButtonEnabled){
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

  question: string = '';

  onChangeQuestion(event:any): void {
    this.question = event.target.value;
  }

  onClickSubmit(){
    if (this.question === ''){
      this.question = 'No answer'
    }
    this.isButtonEnabled = false;
    this.gameService.sendAnswer(this.question)
  }
}
