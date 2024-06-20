import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService} from "../../services/game/game.service";

@Component({
  selector: 'app-invited-user',
  templateUrl: './invited-user.component.html',
  styleUrl: './invited-user.component.scss'
})
export class InvitedUserComponent implements OnInit{
  constructor(private router:Router,private gameService: GameService) {
  }

  ngOnInit(): void {
    if (!this.gameService.connection) {
      this.router.navigateByUrl('/home').then()
    }

    this.gameService.connection?.subscribe((data) => {
      console.log(data)
    })
  }
}
