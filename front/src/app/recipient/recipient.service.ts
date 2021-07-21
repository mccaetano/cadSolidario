import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Recipient } from './recipient';

@Injectable({
  providedIn: 'root'
})
export class RecipientService {

  private apiURL = "/api";

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }

  constructor(private httpClient: HttpClient) { }

  getByFilter(name: string): Observable<Recipient[]> {
    return this.httpClient.get<Recipient[]>(this.apiURL + '/recipient?name='+ name + 
      '&limit=20&skip=1')
    .pipe(
      catchError(this.errorHandler)
    )
  }
  
  getById(id: number): Observable<Recipient> {
    return this.httpClient.get<Recipient>(this.apiURL + '/recipient/'+ id)
    .pipe(
      catchError(this.errorHandler)
    )
  }

  create(recipient: Recipient): Observable<string> {
    return this.httpClient.post<string>(this.apiURL + '/recipient', JSON.stringify(recipient), this.httpOptions)
    .pipe(
      catchError(this.errorHandler)
    )
  }

  update(id: number, recipient: Recipient): Observable<Recipient> {    
    return this.httpClient.put<Recipient>(this.apiURL + '/recipient/' + id, JSON.stringify(recipient), this.httpOptions)
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
    console.log("error: " + errorMessage);
    return throwError(errorMessage);
 }

}
