import {Component, OnInit} from '@angular/core';
import {GameService} from "../services/game/game.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-answer-menu',
  templateUrl: './answer-menu.component.html',
  styleUrl: './answer-menu.component.scss'
})
export class AnswerMenuComponent implements OnInit{
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
