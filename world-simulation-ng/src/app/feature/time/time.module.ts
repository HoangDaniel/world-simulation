import { CommonModule, DatePipe } from "@angular/common";
import { NgModule } from "@angular/core";
import { TimeService } from "./services/time.service";
import { TimeComponent } from "./time.component";
import { HttpClientModule } from "@angular/common/http";

@NgModule({
    imports: [CommonModule, DatePipe, HttpClientModule],
    providers: [TimeService],
    declarations: [TimeComponent],
    exports: [TimeComponent]
})
export class TimeModule { }