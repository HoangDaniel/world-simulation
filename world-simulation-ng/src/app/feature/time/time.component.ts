import { Component, OnInit } from "@angular/core";
import { TimeService } from "./services/time.service";

@Component({
    selector: 'app-time',
    templateUrl: './time.component.html',
})
export class TimeComponent {
    currentTime$ = this.timeService.getTimeUpdates();

    constructor(private timeService: TimeService) { }
}