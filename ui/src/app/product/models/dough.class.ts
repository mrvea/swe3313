import { Product, Productor } from "./product.class";

export interface Dougher extends Productor {
	Type?: DoughTypes
}

export enum DoughTypes { //database?
	THIN = "Thin",
	DEEP = "Deep",
	PAN = "Pan"
}
export class Dough extends Product{
	Type: DoughTypes;
	constructor(options: Dougher = {}){
		super(options);

		this.Type = options["Type"] || DoughTypes.PAN;
	}
}