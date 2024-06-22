import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';


export type CreateRoomResponse = {
  roomId: string;
  extraPrompt: string;
}

export type UpdateRoomRequest = {
  extraPrompt: string;
}

export type UpdateRoomResponse = {
  roomId: string;
  extraPrompt: string;
}


@Injectable({
  providedIn: 'root'
})
export class HttpService {
  private apiURI = "http://ai-phone-alb-345985775.us-east-1.elb.amazonaws.com";
  constructor(private http: HttpClient) {}

  CreateRoom(): Observable<CreateRoomResponse>{
    return this.http.post<CreateRoomResponse>(this.apiURI + '/room',{
      Headers: {
      'Access-Control-Allow-Origin': '*',
      },
    });
  }

  UpdateRoom(roomId: string, extraPrompt: string): Observable<UpdateRoomResponse>{
    let request :UpdateRoomRequest = {
      extraPrompt: extraPrompt,
    }
    return this.http.post<UpdateRoomResponse>(this.apiURI + '/room/'+ roomId, {
      request,
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
    });
  }
}
