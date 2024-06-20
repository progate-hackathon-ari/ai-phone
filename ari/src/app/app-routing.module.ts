import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeScreenComponent } from './home/home-screen/home-screen.component';
import { AdminUserComponent } from './wait-room/admin-user/admin-user.component';
import {QuestionMenuComponent} from "./question-menu/question-menu.component";
import {AnswerMenuComponent} from "./answer-menu/answer-menu.component";

const routes: Routes = [
    {path: '', redirectTo: '/home', pathMatch: 'full'},
    {path: 'home', component: HomeScreenComponent},
    {path: "question",component: QuestionMenuComponent},
    {path: "answer", component:AnswerMenuComponent},
    {path: 'admin', component: AdminUserComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
