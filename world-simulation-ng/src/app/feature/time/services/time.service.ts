import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
    providedIn: 'root'
})
export class TimeService {
    private timeSocket: WebSocketSubject<Date>;
    private apiUrl = 'http://localhost:8080'

    constructor(private http: HttpClient) {
        this.timeSocket = webSocket('ws://localhost:8080/ws');
    }

    public getTimeUpdates(): Observable<Date> {
        return this.timeSocket.asObservable();
    }

    setSpeed(speed: number): Observable<any> {
        return this.http.post(`${this.apiUrl}/setSpeed`, { speed })
    }

    doAction(action: string): Observable<any> {
        const url = `${this.apiUrl}/${action}`;
        return this.http.get(url);
    }
}