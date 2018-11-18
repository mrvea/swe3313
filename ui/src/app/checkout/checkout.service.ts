import { Injectable } from '@angular/core';
import { Observable, interval, of } from 'rxjs';
import { delay } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class CheckoutService {

  constructor() { }


  processCard(prder): Observable<boolean> {
    return of(true).pipe(delay(1000));
  }

  processCheck(order): Observable<boolean> {
    return of(true).pipe(delay(1000));
  }
}
