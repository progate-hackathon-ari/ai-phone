import {Component, OnInit} from '@angular/core';
import {GameService} from "../../services/game/game.service";

@Component({
  selector: 'app-home-screen',
  templateUrl: './home-screen.component.html',
  styleUrl: './home-screen.component.scss'
})
export class HomeScreenComponent implements OnInit{
  constructor(private gameService: GameService) {
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
}
