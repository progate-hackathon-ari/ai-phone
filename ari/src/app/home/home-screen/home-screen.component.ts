import { Component } from '@angular/core';
import { HttpService } from '../../serivce/http.service';

@Component({
  selector: 'app-home-screen',
  templateUrl: './home-screen.component.html',
  styleUrl: './home-screen.component.scss'
})
export class HomeScreenComponent {
  constructor(private http: HttpService){}

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
