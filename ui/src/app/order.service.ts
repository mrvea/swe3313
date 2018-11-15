import { Injectable } from '@angular/core';
import { Product } from './product/models';
import { Subject, BehaviorSubject, Observable } from 'rxjs';
import { tap, take, catchError, shareReplay, scan, map } from 'rxjs/operators';

export enum PayloadAction {
	add = "add",
	updateToppings = "update-toppings",
	update = "update",
	delete = "delete"
}

export interface Payloader {
	action?: PayloadAction,
	product?: Product[],
	key?: string,
	index?: number,
}
export class SendPayload implements Payloader {
	index: number;
	key: string;
	action: PayloadAction;
	product: Product[];
	constructor(options: Payloader = {}){
		this.index = options.index || -1;
		this.key = options.key || "";
		this.action = options.action || PayloadAction.add;
		this.product = options.product || [];
	}
}

@Injectable({
  providedIn: 'root'
})
export class OrderService {

	private _order$ = new Subject<SendPayload>();
 	order$ = this._order$.pipe(
 		scan((acc: SendPayload, cur: SendPayload) => {
 			switch (cur.action) {
 				case PayloadAction.add:
 					console.log(acc, cur);
 					return new SendPayload({product: cur.product.concat(acc.product)});
 					
 				case PayloadAction.update:
 				console.log(acc, cur);
 					var send = new SendPayload(acc);
 					send.product[cur.index][cur.key] = cur.product[cur.key]
 					return send;
 					// acc.product[cur.index][cur.key] = cur.product;
  				case PayloadAction.updateToppings:
 				console.log(acc, cur);
 					var send = new SendPayload(acc);
 					if(cur.index < 0){
 						// send.product[cur.index*-1].Toppings =
 					}
 					send.product[cur.index]["Toppings"] = send.product[cur.index]["Toppings"].concat(cur.product);
 					console.log(send);
 					return send;
 					// acc.product[cur.index][cur.key].push(cur.product);
				case PayloadAction.delete:
				console.log(acc, cur);
					acc.product.splice(cur.index, 1);
				break;
 				default:
 					// code...
 					console.log("could not find acction: ", cur.action);
 					break;
 			}
 			
 			return new SendPayload(acc);
 		}, new SendPayload()),
 		shareReplay(1)
 	);
	public itemCount$ = this.order$.pipe(
			map((o: SendPayload) => o.product.length),
			shareReplay(1)
		);

  constructor() { }

  addToOrder(p: Product): void {
  	this._order$.next(new SendPayload({product: [p], action: PayloadAction.add}));
  }
  addToppingTo(index, p: Product){
  	this._order$.next(new SendPayload({index: index, product: [p], action: PayloadAction.updateToppings}))
  }
  remove(index){
  	this._order$.next(new SendPayload({product: [], action: PayloadAction.delete, index: index}));
  }
  removeTopping(index){
  	this._order$.next(new SendPayload({product: [], action: PayloadAction.updateToppings, index: index * -1}))
  }
}
