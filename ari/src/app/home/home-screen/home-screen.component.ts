import {Component, OnInit} from '@angular/core';
import {GameService} from "../../services/game/game.service";
import { HttpService } from '../../services/http/http.service';
import {Router} from "@angular/router";

@Component({
  selector: 'app-home-screen',
  templateUrl: './home-screen.component.html',
  styleUrl: './home-screen.component.scss'
})
export class HomeScreenComponent implements OnInit{
  constructor(
    private router: Router,
    private gameService: GameService,
    private http: HttpService,
  ) {
  }

  idValue: string = "";
  nameValue: string = "";

  ngOnInit() {
    this.gameService.connect();
  }

  onChangeIdInput(event: any) {
    this.idValue = event.target.value;
  }

  onChangeNameInput(event: any) {
    this.nameValue = event.target.value;
  }

  onClickJoinRoom() {
    this.gameService.sendJoin(this.idValue,this.nameValue);
    this.router.navigateByUrl("/invited").then();
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
