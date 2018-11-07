import { Product, Productor } from "./product.class";

export interface Saucerizer extends Productor {
	Type?: SauceTypes;
}

export enum SauceTypes { //database?
	RED = "Tomato",
	WHITE = "Alfredo",
	Pink = "red/white"
}
export class Sauce extends Product implements Saucerizer {
	Type: SauceTypes;
	constructor(options: Saucerizer = {}){
		super(options);

		this.Type = options.Type || SauceTypes.RED;
	}
}