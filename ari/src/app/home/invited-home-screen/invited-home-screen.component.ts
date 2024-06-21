import {Component, OnDestroy, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import { GameService, dataSubscribe } from '../../services/game/game.service';
import {Observable, Subscription} from "rxjs";
interface PlayerData {
  connection_id: string
  index: number
  room_id: string
  username: string
}
@Component({
  selector: 'app-invited-home-screen',
  templateUrl: './invited-home-screen.component.html',
  styleUrl: './invited-home-screen.component.scss'
})
export class InvitedHomeScreenComponent implements OnInit, OnDestroy {
  constructor(
      private router: Router,
      private route: ActivatedRoute,
      private gameService: GameService,
      private dataSubscribe: dataSubscribe
  ) {
  }
  roomId: string = "";
  dsub: Observable<any> | undefined;
  Subs: Subscription | undefined;
  players: PlayerData[] = [];

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.roomId = params["roomId"];
    })

    this.dsub = this.dataSubscribe.subscribe();

    this.Subs = this.dsub.subscribe(data => {
      const json = JSON.parse(data)
      this.players = json.players
    })
  }

  ngOnDestroy(): void {
    if (this.Subs) {
      this.Subs.unsubscribe();
    }
  }

  nameValue: string = "";

  onChangeNameInput(event: any) {
    this.nameValue = event.target.value;
  }

  onClickJoinRoom() {
    this.gameService.sendJoin(this.roomId,this.nameValue);
    this.router.navigateByUrl("/invited").then();
  }
}
