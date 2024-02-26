import { CommonModule, DatePipe } from "@angular/common";
import { NgModule } from "@angular/core";
import { TimeService } from "./services/time.service";
import { TimeComponent } from "./time.component";

@NgModule({
    imports: [CommonModule, DatePipe],
    providers: [TimeService],
    declarations: [TimeComponent],
    exports: [TimeComponent]
})
export class TimeModule { }