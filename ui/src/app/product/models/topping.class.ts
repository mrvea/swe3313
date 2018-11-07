import { Product, Productor } from './product.class';

export interface Toppinger extends Productor {
	Type?: ToppingTypes;
}

export enum ToppingTypes {//database?
	MEAT = "Meat",
	VEG = "Vegetable",
	CHEESE = "Cheese",
	UNKNOWN = "Unknown"
}
export class Topping extends Product {
	Type: ToppingTypes;
	constructor(options: Toppinger = {}){
		super(options);

		this.Type = options.Type || ToppingTypes.UNKNOWN;
	}
}