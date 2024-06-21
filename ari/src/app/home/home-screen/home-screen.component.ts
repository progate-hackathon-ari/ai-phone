import {Component} from '@angular/core';
import {GameService} from "../../services/game/game.service";
import { HttpService } from '../../services/http/http.service';
import {Router} from "@angular/router";

@Component({
  selector: 'app-home-screen',
  templateUrl: './home-screen.component.html',
  styleUrl: './home-screen.component.scss'
})
export class HomeScreenComponent{
  constructor(
    private router: Router,
    private gameService: GameService,
    private http: HttpService,
  ) {
  }
  nameValue: string = "";

  onChangeNameInput(event: any) {
    this.nameValue = event.target.value;
  }

  onClickCreateRoom() {
    if (this.nameValue === "") {
      alert("Please enter your name");
    }else{
      let room = this.http.CreateRoom();

      room.subscribe(data => {
        this.gameService.sendJoin(data.roomId,this.nameValue);
        this.router.navigateByUrl("/admin").then();
      })
    }
  }
}
