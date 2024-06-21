import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService} from "../../services/game/game.service";

interface PlayerData {
  connection_id: string
  index: number
  room_id: string
  username: string
}

@Component({
  selector: 'app-admin-user',
  templateUrl: './admin-user.component.html',
  styleUrl: './admin-user.component.scss'
})
export class AdminUserComponent implements OnInit{
  constructor(private router:Router,private gameService: GameService) {
  }

  players: PlayerData[] = []

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    this.gameService.connect().subscribe(this.subscribeHandler)
  }

  subscribeHandler(data: string){
    const json = JSON.parse(data)
    this.players = json.players
    console.log(this.players)
  }

  onClickStart(){
    this.gameService.sendReady()
    this.router.navigateByUrl("/question").then()
  }
}
