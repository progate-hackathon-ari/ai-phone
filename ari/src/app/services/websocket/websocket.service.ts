import {Observable, Observer, Subject} from "rxjs";
import { webSocket } from "rxjs/webSocket";
import {Injectable} from "@angular/core";

@Injectable()
export class WebSocketService {

  connect(url: string): Subject<string> {
    return webSocket({
      url: url,
      deserializer: (e: MessageEvent) => e.data,
    })
  }
}
