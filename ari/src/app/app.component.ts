import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import { GameService } from './services/game/game.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit{
  title = 'ari';

  constructor(
    private router: Router,
    private gameService: GameService
  ) {
  }

  ngOnInit(): void {
    this.gameService.connect();
  }
}
