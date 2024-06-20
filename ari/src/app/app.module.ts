import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeScreenComponent } from './home/home-screen/home-screen.component';
import { AdminUserComponent } from './wait-room/admin-user/admin-user.component';
import { QuestionMenuComponent } from './question-menu/question-menu.component';
import { AnswerMenuComponent } from './answer-menu/answer-menu.component';
import { ResultMenuComponent } from './result-menu/result-menu.component';
import {WebSocketService} from "./services/websocket/websocket.service";


@NgModule({
  declarations: [
    AppComponent,
    HomeScreenComponent,
    AdminUserComponent,
    QuestionMenuComponent,
    AnswerMenuComponent,
    ResultMenuComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [
    WebSocketService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
