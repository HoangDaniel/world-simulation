import { Observable } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

export class TimeService {
    private timeSocket: WebSocketSubject<Date>;

    constructor() {
        this.timeSocket = webSocket('ws://localhost:8080/ws');
    }

    public getTimeUpdates(): Observable<Date> {
        return this.timeSocket.asObservable();
    }
}