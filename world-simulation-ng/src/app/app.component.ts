import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { TimeModule } from './feature/time/time.module';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, TimeModule],
  templateUrl: './app.component.html',
})
export class AppComponent {
}
