import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {Observable, Subscription} from "rxjs";
import { GameService, dataSubscribe } from '../../services/game/game.service';
import {HttpService} from "../../services/http/http.service";

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
export class AdminUserComponent implements OnInit , OnDestroy{
  constructor(private router: Router, private gameService: GameService, private dataSubscribe: dataSubscribe,private http: HttpService) {
}

  players: PlayerData[] = []
  subscription: Subscription | undefined
  dsub: Observable<any> | undefined;
  Subs: Subscription | undefined;
  roomId: string | undefined = this.gameService.roomId

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    this.dsub = this.dataSubscribe.subscribe();

    this.Subs = this.dsub.subscribe(data => {
        const json = JSON.parse(data)
        this.players = json.players
        console.log(this.players)
    })

  }

  ngOnDestroy(): void {
    if (this.Subs) {
      this.Subs.unsubscribe();
    }
  }

  selectedOption: string = ''; // デフォルト値を設定

  selectOption(option: string): void {
    this.selectedOption = option;
  }

copyUrl(){
    navigator.clipboard.writeText(`${document.location.origin}/invited-home?roomId=${this.roomId}`).then();
    alert("Copied the URL");
}

  onClickStart(){
    if (this.roomId === undefined) return
    const result = this.http.UpdateRoom(this.roomId, this.selectedOption)

    result.subscribe(data => {
        this.gameService.sendReady()
        this.gameService.isAdmin = true
    })
        this.router.navigateByUrl("/countdown").then()
  }
}