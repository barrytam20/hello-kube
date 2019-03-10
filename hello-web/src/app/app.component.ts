import { HelloService } from './hello.service';
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  greeting: string;
  weather: string;
  constructor(private helloService: HelloService) { }

  greetName(name: string): void {
    if (name) {
      this.helloService.getGreeting(name).subscribe(message => {
        this.greeting = message.greeting;
      });
    }
  }

  displayWeather(zip: string): void {
    if (zip) {
      this.helloService.getWeather(zip).subscribe(forecast => {
        if (forecast.description) {
          this.weather = `Currently, it is ${forecast.description} in ${forecast.city}, ${forecast.state}`;
        }
      });
    }
  }
}
