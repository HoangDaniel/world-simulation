import { Component, OnInit } from "@angular/core";
import { TimeService } from "./services/time.service";
import { BehaviorSubject } from "rxjs";

@Component({
    selector: 'app-time',
    templateUrl: './time.component.html',
})
export class TimeComponent {
    currentTime$ = this.timeService.getTimeUpdates();
    paused$ = new BehaviorSubject<boolean>(false)

    constructor(private timeService: TimeService) { }

    pause() {
        this.timeService.doAction('pause').subscribe(() => {
            this.paused$.next(true)
        })
    }

    resume() {
        this.timeService.doAction('resume').subscribe(() => {
            this.paused$.next(false)
        })
    }

    setSpeed(event: Event) {
        const selectElement = event.target as HTMLSelectElement;
        const speedValue = Number(selectElement.value);
        this.timeService.setSpeed(speedValue).subscribe((res) => {
            console.log(res)
        });
    }
}