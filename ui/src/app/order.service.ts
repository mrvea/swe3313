import { Injectable } from '@angular/core';
import { Product } from './product/models';
import { Subject, BehaviorSubject, Observable } from 'rxjs';
import { tap, take, catchError, shareReplay, scan, map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class OrderService {

	private _order$: Subject<Product[]> = new BehaviorSubject<Product[]>([]);
	public order$ = this._order$.pipe(
		scan((acc: Product[], cur: Product[]) => acc.concat(cur)),
		shareReplay(1));
	public itemCount$ = this.order$.pipe(
			map(o => o.length),
			shareReplay(1)
		);

  constructor() { }

  addToOrder(p: Product): void {
  	this._order$.next([p]);
  }
}
