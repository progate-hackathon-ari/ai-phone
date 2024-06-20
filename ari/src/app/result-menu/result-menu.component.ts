import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService} from "../services/game/game.service";

@Component({
  selector: 'app-result-menu',
  templateUrl: './result-menu.component.html',
  styleUrl: './result-menu.component.scss'
})
export class ResultMenuComponent implements OnInit{
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
