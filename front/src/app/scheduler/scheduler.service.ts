import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { Scheduler } from './scheduler';
import { catchError } from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class SchedulerService {

  private apiURL = "/api";

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }

  constructor(private httpClient: HttpClient) { }

  getAll(eventDate: Date, name: string, page: number): Observable<Scheduler[]> {
    return this.httpClient.get<Scheduler[]>(this.apiURL + '/calendar?eventDate='+
      eventDate.toISOString().split('T')[0] + "&name=" + name + "&limit=20&skip=" + page)
    .pipe(
      catchError(this.errorHandler)
    )
  }

  getEventDates(): Observable<Scheduler[]> {
    return this.httpClient.get<Scheduler[]>(this.apiURL + '/calendar/event-dates')
    .pipe(
      catchError(this.errorHandler)
    )
  }

  create(scheduler: Scheduler): Observable<Scheduler> {
    if (scheduler.eventDate == null) { scheduler.eventDate = new Date("1900-01-01")}

    return this.httpClient.post<Scheduler>(this.apiURL + '/calendar', JSON.stringify(scheduler), this.httpOptions)
    .pipe(
      catchError(this.errorHandler)
    )
  }

  update(id?: number, scheduler?: Scheduler): Observable<Scheduler> {
    console.log('update: ' + JSON.stringify(scheduler));
    return this.httpClient.put<Scheduler>(this.apiURL + '/calendar/' + id, JSON.stringify(scheduler), this.httpOptions)
    .pipe(
      catchError(this.errorHandler)
    )
  }

  errorHandler(error: any) {
    let errorMessage = '';
    if(error.error instanceof ErrorEvent) {
      errorMessage = error.error.message;
    } else {
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    return throwError(errorMessage);
 }
}
