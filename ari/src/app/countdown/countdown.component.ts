import {Component, OnInit} from '@angular/core';
import {GameService} from "../services/game/game.service";
import {Router} from "@angular/router";
interface CountdownData {
  is_done: boolean
  count: number
}

@Component({
  selector: 'app-countdown',
  templateUrl: './countdown.component.html',
  styleUrl: './countdown.component.scss'
})
export class CountdownComponent implements OnInit{
  countNumber: number | undefined
  constructor(private router: Router, private gameService: GameService) {
  }

  ngOnInit(): void {
    // this.gameService.getSubscribe().subscribe((data)=>{
    //   let countdownData: CountdownData = JSON.parse(data);
    //   this.countNumber = countdownData.count;
    //   if (countdownData.is_done) {
    //     this.router.navigateByUrl('/question').then();
    //   }
    // });

    this.gameService.removeSubscribe()
  }
}
