import { Component, ViewChild } from '@angular/core';
import { Subject } from 'rxjs';
import { scan, shareReplay } from 'rxjs/operators';

import { OrderService, SendPayload } from './order.service';
import { Pizza, Dough, Topping, Sauce, DoughType, DoughSize, ToppingType, SauceType, Product } from './product/models';

const MENU_LIST = [
	{
		Name: "item",
		Path: ""
	},
	{
		Name: "item2",
		Path: ""
	},
	{
		Name: "item3",
		Path: ""
	}
];
@Component({
  selector: 'app-home',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
	MenuList: any[] = [];
 	Title = "Geb's Pizza";
 	order: Product[] = [];
 	doughType = DoughType;
 	doughSize = DoughSize;
 	sauceType = SauceType;

 	@ViewChild("orderSidenav") cart;
 	constructor(
 		protected os: OrderService
 	){
 		this.MenuList = MENU_LIST;

 		this.os.order$.subscribe((o: SendPayload) => {
 			console.log(o);
 			console.log(this.cart);
 			if((!this.cart.opened && o.product.length > 0)|| 
 				this.cart.opened && o.product.length == 0){
 				this.cart.toggle();
 			}
 			this.order = o.product;
 		})
 	}

 	UpdateWidth(){

 	}
	GetAnimation(o){

	}
	getDoughSizes(): any[]{
		return Dough.sizes();
	}

	getDoughTypes(): string[]{
		return Dough.types();
	}

	getSauceTypes(): string[]{
		return Sauce.types();
	}

	remove(product, index){
		this.os.remove(index);
	}

	removeTopping(i){
		this.os.removeTopping(i);
	}
}
