import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";

@Component({
  selector: 'app-none',
  templateUrl: './none.component.html',
  styleUrl: './none.component.scss'
})
export class NoneComponent implements OnInit{
  constructor(private router: Router, ) {
  }

  ngOnInit(): void {
    this.router.navigateByUrl('/answer').then()
  }
}
