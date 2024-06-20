import { Injectable } from '@angular/core';
import {Subject} from "rxjs";
import {webSocket} from "rxjs/webSocket";

enum EventType {
  EventJoin = 'join',
  EventAnswer = 'answer',
  EventReady = 'ready',
  EventNext = 'next',
  EventCountDown = 'countdown',
}


type MessageTemplate = {
  event: EventType,
  roomId: string,
  data?: string,
}

@Injectable({
  providedIn: 'root'
})
export class GameService {
  connection: Subject<string> | undefined

  connect(){
    // TODO: envからendpointを取るようにする
    this.connection = webSocket({
      url: `ws://localhost:8080/game`,
      deserializer: (e: MessageEvent) => e.data,
    })
    return this.connection
  }

  sendJoin(roomId: string, name: string): void {
    let data = JSON.stringify({
      name: name,
    })

    let message: MessageTemplate = {
      event: EventType.EventJoin,
      roomId: roomId,
      data: btoa(data),
    }

    this.sendData(JSON.stringify(message))
  }

  sendAnswer(roomId: string, answer: string): void {
    let data = JSON.stringify({
      answer: answer,
    })

    let message: MessageTemplate = {
      event: EventType.EventAnswer,
      roomId: roomId,
      data: btoa(data),
    }

    this.sendData(JSON.stringify(message))
  }

  sendReady(roomId: string): void {
    let message: MessageTemplate = {
      event: EventType.EventReady,
      roomId: roomId,
    }

    this.sendData(JSON.stringify(message))
  }

  sendNext(roomId: string): void {
    let message: MessageTemplate = {
      event: EventType.EventNext,
      roomId: roomId,
    }

    this.sendData(JSON.stringify(message))
  }

  sendCountDown(roomId: string, count: number): void {
    let data = JSON.stringify({
      count: count,
    })

    let message: MessageTemplate = {
      event: EventType.EventCountDown,
      roomId: roomId,
      data: btoa(data),
    }

    this.sendData(JSON.stringify(message))
  }

  // backendがjsonとしてparseできるようにbase64でエンコード
  sendData(data: string): void {
    if (!this.connection) {
      throw new Error('connection is not initialized')
    }
    this.connection.next(btoa(data))
  }
}
