import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {Observable, Subscription} from "rxjs";
import { GameService, dataSubscribe } from '../../services/game/game.service';

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
export class AdminUserComponent implements OnInit {
  constructor(private router: Router, private gameService: GameService, private dataSubs: dataSubscribe) {
    
  }

  players: PlayerData[] = []
  subscription: Subscription | undefined
  dsub: Observable<any> | undefined;
  Subs: Subscription | undefined;

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    // this.subscription = this.gameService.getSubscribe().subscribe((data) => {
    //     const json = JSON.parse(data)
    //     this.players = json.players
    //     console.log(this.players)
    //   })

    this.dsub = this.dataSubs.subscribe();

    this.Subs = this.dsub.subscribe(data => {
        const json = JSON.parse(data)
        this.players = json.players
        console.log(this.players)
    })
    
  }

  onClickStart() {
    this.gameService.sendReady()
    this.router.navigateByUrl("/countdown").then(()=>{
      setTimeout(()=>{
        this.subscription?.unsubscribe()
      },1000)
    })
  }
}
