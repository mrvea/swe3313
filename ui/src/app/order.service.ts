import { Injectable } from '@angular/core';
import { Product, Topping } from './product/models';
import { Subject, BehaviorSubject, Observable } from 'rxjs';
import { tap, take, catchError, shareReplay, scan, map, reduce } from 'rxjs/operators';
import { User } from './user/user.class';
import { HttpClient } from '@angular/common/http';
import { Order } from './order.class';

export enum PayloadAction {
    add = 'add',
    updateToppings = 'update-toppings',
    update = 'update',
    delete = 'delete',
    reset = 'reset'
}

export interface Payloader {
    action?: PayloadAction;
    product?: Product[];
    key?: string;
    index?: number;
}
export class SendPayload implements Payloader {
    index: number;
    key: string;
    action: PayloadAction;
    product: Product[];
    constructor(options: Payloader = {}) {
        this.index = options.index || -1;
        this.key = options.key || '';
        this.action = options.action || PayloadAction.add;
        this.product = options.product || [];
    }
}

@Injectable({
  providedIn: 'root'
})
export class OrderService {

    private _order$ = new Subject<SendPayload>();
    public order$ = this._order$.pipe(
         scan((acc: SendPayload, cur: SendPayload) => {
             let send: SendPayload;
             switch (cur.action) {
                 case PayloadAction.add:
                     console.log(acc, cur);
                     return new SendPayload({product: cur.product.concat(acc.product)});

                 case PayloadAction.update:
                 console.log(acc, cur);
                     send = new SendPayload(acc);
                     send.product[cur.index][cur.key] = cur.product[cur.key];
                     return send;
                     // acc.product[cur.index][cur.key] = cur.product;
                  case PayloadAction.updateToppings:
                 console.log(acc, cur);
                     send = new SendPayload(acc);
                     if (cur.index < 0) {
                         // send.product[cur.index*-1].Toppings =
                     }
                     send.product[cur.index]['Toppings'] = send.product[cur.index]['Toppings'].concat(cur.product);
                     console.log(send);
                     return send;
                     // acc.product[cur.index][cur.key].push(cur.product);
                case PayloadAction.delete:
                console.log(acc, cur);
                    acc.product.splice(cur.index, 1);
                break;
                case PayloadAction.reset:
                     return new SendPayload();
                 default:
                     // code...
                     console.log('could not find acction: ', cur.action);
                     break;
             }

             return new SendPayload(acc);
         }, new SendPayload()),
         shareReplay(2)
     );

    public itemCount$ = this.order$.pipe(
            map((o: SendPayload) => o.product.length),
            shareReplay(2)
        );

    public totalPrice$ = this.order$.pipe(
        map((o: SendPayload) => o.product.reduce((a, b) => a + Product.getPrice(b), 0)),
        reduce((a, b) => a + b, 0),
        tap(price => console.log(price)),
        shareReplay(2)
    );
    private _customer$ = new Subject<User>();
    customer$ = this._customer$.pipe(shareReplay(1));
    private _customer: User = null;

      constructor(
          private http: HttpClient
      ) { }

      addToOrder(p: Product): void {
          this._order$.next(new SendPayload({product: [p], action: PayloadAction.add}));
      }
      addToppingTo(index, p: Product) {
          this._order$.next(new SendPayload({index: index, product: [p], action: PayloadAction.updateToppings}));
    }
      remove(index) {
          this._order$.next(new SendPayload({product: [], action: PayloadAction.delete, index: index}));
      }
      removeTopping(pIndex, t: Topping) {
          this._order$.next(new SendPayload({product: [t], action: PayloadAction.updateToppings, index: pIndex * 1}));
    }

    setCustomer(c: User): void {
        this._customer$.next(c);
    }

    reset(): void {
        this._customer$.next(null);
        this._order$.next(new SendPayload({action: PayloadAction.reset}));
    }

    save(o: Order) {
        return this.http.post('/api/orders', o);
    }

}
