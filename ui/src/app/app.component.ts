import { Component, ViewChild } from '@angular/core';
import { Subject } from 'rxjs';
import { scan, shareReplay, tap } from 'rxjs/operators';

import { OrderService, SendPayload } from './order.service';
import { Pizza, Dough, Topping, Sauce, DoughType, DoughSize, ToppingType, SauceType, Product } from './product/models';
import { User } from './user/user.class';

const MENU_LIST = [
    {
        Name: 'item',
        Path: ''
    },
    {
        Name: 'item2',
        Path: ''
    },
    {
        Name: 'item3',
        Path: ''
    }
];
@Component({
  selector: 'app-home',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
    menuList: any[] = [];
// tslint:disable-next-line:quotemark
    title = "Geb's Pizza";
    order: Product[] = [];
    customer: User;
    doughType = DoughType;
    doughSize = DoughSize;
    sauceType = SauceType;
    price = 0;
    @ViewChild('orderSidenav') cart;
    constructor(
        protected os: OrderService
    ) {
        this.menuList = MENU_LIST;

        this.os.order$.pipe(
            tap((o: SendPayload) => this.price = Product.getTotal(o.product))
        ).subscribe((o: SendPayload) => {
            console.log(o);
            console.log(this.cart);
            if ((!this.cart.opened && o.product.length > 0) ||
                this.cart.opened && o.product.length === 0) {
                this.cart.toggle();
            }
            this.order = o.product;
        });
        this.os.totalPrice$.subscribe(price => {
            console.log(price);
            this.price = price;
        });

        this.os.customer$.subscribe(c => {
            console.log(c);
            this.customer = c;
        });
    }

    updateWidth() {

    }
    getAnimation(o) {

    }
    getDoughSizes(): string[] {
        return Dough.sizes();
    }

    getDoughTypes(): string[] {
        return Dough.types();
    }

    getSauceTypes(): string[] {
        return Sauce.types();
    }

    remove(product, index) {
        this.os.remove(index);
    }

    removeTopping(i, t: Topping) {
        this.os.removeTopping(i, t);
    }
}
