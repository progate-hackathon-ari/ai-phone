import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {GameService} from "../../services/game/game.service";

@Component({
  selector: 'app-admin-user',
  templateUrl: './admin-user.component.html',
  styleUrl: './admin-user.component.scss'
})
export class AdminUserComponent implements OnInit{
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

  selectedOption: string = 'option1'; // デフォルト値を設定

  selectOption(option: string): void {
    this.selectedOption = option;
  }
}
