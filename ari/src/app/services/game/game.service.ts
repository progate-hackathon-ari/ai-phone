import {Injectable} from '@angular/core';
import {webSocket, WebSocketSubject} from "rxjs/webSocket";
import {Observable} from "rxjs";

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
  connection: WebSocketSubject<string> | undefined
  roomId: string | undefined

  connect(): WebSocketSubject<string>{
    // TODO: envからendpointを取るようにする
    if (!this.connection) {
      this.connection = webSocket({
        url: `ws://localhost:8080/game`,
        deserializer: (e: MessageEvent) => e.data,
      })
    }
    return this.connection
  }

  getSubscribe(): Observable<string> {
    if (!this.connection) {
      throw new Error('connection is not initialized')
    }

    return new generateMultiplexWebSocket(this.connection).getObservable()
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
    this.roomId = roomId
  }

  sendAnswer(answer: string): void {
    if (!this.roomId) {
      throw new Error('roomId is not initialized')
    }

    let data = JSON.stringify({
      answer: answer,
    })

    let message: MessageTemplate = {
      event: EventType.EventAnswer,
      roomId: this.roomId,
      data: btoa(data),
    }

    this.sendData(JSON.stringify(message))
  }

  sendReady(): void {
    if (!this.roomId) {
      throw new Error('roomId is not initialized')
    }

    let message: MessageTemplate = {
      event: EventType.EventReady,
      roomId: this.roomId,
    }

    this.sendData(JSON.stringify(message))
  }

  sendNext(): void {
    if (!this.roomId) {
      throw new Error('roomId is not initialized')
    }

    let message: MessageTemplate = {
      event: EventType.EventNext,
      roomId: this.roomId,
    }

    this.sendData(JSON.stringify(message))
  }

  sendCountDown(count: number): void {
    if (!this.roomId) {
      throw new Error('roomId is not initialized')
    }

    let data = JSON.stringify({
      count: count,
    })

    let message: MessageTemplate = {
      event: EventType.EventCountDown,
      roomId: this.roomId,
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

class generateMultiplexWebSocket {
  constructor(private connection: WebSocketSubject<string>) {}


  getObservable():  Observable<string>{
    return this.connection.multiplex(
      () => (""),
      () => (""),
      () => true
    )
  }
}
