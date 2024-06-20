import { Injectable } from '@angular/core';
import {Subject} from "rxjs";
import {WebSocketService} from "../websocket/websocket.service";

@Injectable({
  providedIn: 'root'
})
export class GameService {
  connection: Subject<MessageEvent<string>> | undefined

  constructor(private ws: WebSocketService) {
  }

  connect(roomId: string){
    if (!this.ws) throw new Error('No websocket service');
    this.connection = this.ws.connect(`ws://localhost:8080/${roomId}`)
    return this.connection
  }

  sendData(data: string): void {
    if (!this.ws) throw new Error('No connection');
    this.connection?.next(new MessageEvent('message', {data}));
  }
}
