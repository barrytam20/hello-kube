import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HelloService {
  private httpOptions = {
    headers: new HttpHeaders({
      accept: 'application/json;v=1'
    })
  };
  constructor(private http: HttpClient) { }

  getGreeting(name: string): Observable<Greeting> {
    return this.http.get<Greeting>(`/hello?name=${name}`, this.httpOptions);
  }

  getWeather(zip: string): Observable<Weather> {
    return this.http.get<Weather>(`/weather/${zip}`, this.httpOptions);
  }
}

export interface Greeting {
  greeting: string;
}

export interface Weather {
  description: string;
  iconLink: string;
  country: string;
  state: string;
  city: string;
}
