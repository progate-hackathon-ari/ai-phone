import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {dataSubscribe, GameService} from "../services/game/game.service";
import {Observable, Subscription} from "rxjs";

@Component({
  selector: 'app-question-menu',
  templateUrl: './question-menu.component.html',
  styleUrl: './question-menu.component.scss'
})
export class QuestionMenuComponent implements OnInit{
  constructor(private router:Router,private gameService: GameService, private dataSubs: dataSubscribe ) {
  }
  questionValue: string = "";
  dsub: Observable<any> | undefined;
  Subs: Subscription | undefined;

  ngOnInit(): void {
      this.dsub = this.dataSubs.subscribe();

      this.Subs = this.dsub?.subscribe(data => {
          const json = JSON.parse(data)
          if (this.gameService.isAdmin && json.type === "true") {
              this.router.navigateByUrl('/answer').then()
          } else if (!this.gameService.isAdmin && json.type === "false") {
                this.router.navigateByUrl('/result').then()
          }
      })
  }

  onChangeQuestionInput(event: any) {
    this.questionValue = event.target.value;
  }

  onSendButton(){
    this.gameService.sendAnswer(this.questionValue);
  }
}
