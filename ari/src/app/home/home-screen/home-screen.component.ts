import {Component, OnInit} from '@angular/core';
import {GameService} from "../../services/game/game.service";
import { HttpService } from '../../services/http/http.service';

@Component({
  selector: 'app-home-screen',
  templateUrl: './home-screen.component.html',
  styleUrl: './home-screen.component.scss'
})
export class HomeScreenComponent implements OnInit{
  constructor(
    private gameService: GameService,
    private http: HttpService,
  ) {
  }

  send(data: string): void {
    this.gameService.sendData(
      data
    );
  }

  ngOnInit() {
    console.log("home screen init");
    this.gameService.connect().subscribe(data => {
      console.log(data);
    })
    console.log("send join");
    this.gameService.sendJoin("019034c5-2936-73c7-901c-3ee4480729c9","hoge");
  } 
  
  creatAndJoinRoom() {
    let room = this.http.CreateRoom();

    room.subscribe(data => {
      this.joinRoom(data.roomId);
    })
  }

  joinRoom(roomId: string) {
    console.log(roomId);
  }
}
