import {Component, OnInit} from '@angular/core';
import {GameService} from "../../services/game/game.service";
import { HttpService } from '../../serivce/http.service';

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
    console.log('Home screen')
    this.gameService.connect('/').subscribe(data => {
      console.log('Home screen')
      console.log(data);
    })

    this.send('Hello')
    console.log('Home screen')
  } 
  creatAndJoinRoom() {
    let room = this.http.CreateRoom();

    room.subscribe(data => {
      this.joinRoom(data.roomId);
    })
  }

  joinRoom(roomId: string) {
    // このタイミングでws確立させて join room を行う
    console.log(roomId);
  }
}
