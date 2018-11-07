import { Product, Productor } from './product.class';
import { Dough, Dougher } from './dough.class';
import { Sauce, Saucerizer } from './sauce.class';
import { Topping, Toppinger } from './topping.class';

export interface Pizzariler extends Productor {
	Dough?: Dougher;
	Sauce?: Saucerizer;
	Toppings?: Toppinger[];
}

export class Pizza extends Product implements Pizzariler{
	Dough: Dough;
	Sauce: Sauce;
	Topppings: Topping[];
	constructor(options: Pizzariler = {}){
		super(options);

		this.Dough = options.Dough? new Dough(options.Dough): null;
		this.Sauce = options.Sauce? new Sauce(options.Sauce): null;
		this.Topppings = options.Toppings? this.setToppings(options.Toppings) :[];
	}

	setToppings(toppings: any[]){
		if(toppings.length == 0){
			return toppings;
		}
		return toppings.map(topping => {
			return new Topping(topping);
		});
	}
}