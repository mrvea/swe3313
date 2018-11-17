import { Product, Productor } from "./product.class";

export interface Saucerizer extends Productor {
	Type?: SauceType;
}

export enum SauceType { //database?
	RED = "Tomato",
	WHITE = "Alfredo",
	Pink = "red/white"
}
export class Sauce extends Product implements Saucerizer {
	Type: SauceType;
	constructor(options: Saucerizer = {}){
		super(options);

		this.Type = options.Type || SauceType.RED;
		this.Price = .99;T
	}
	static types(): string[]{
		return Object.keys(SauceType);
	}
	copy(): Sauce{
		return new Sauce(this);
	}
}