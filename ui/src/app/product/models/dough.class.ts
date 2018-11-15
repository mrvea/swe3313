import { Product, Productor } from "./product.class";

export interface Dougher extends Productor {
	Type?: DoughType
	Size?: DoughSize
	Prices?: DoughPrices
}

export enum DoughSize {
	SMALL = "Small",
	MEDIUM = "Medium",
	LARGE = "Large"
}

export enum DoughType { //database?
	THIN = "Thin",
	DEEP = "Deep",
	PAN = "Pan"
}
export interface DoughPrices {
	Small: {
			Thin: number,
			Deep: number,
			Pan: number
		};
	Medium: {
			Thin: number,
			Deep: number,
			Pan: number
		};
	Large: {
			Thin: number,
			Deep: number,
			Pan: number
		}
}
export class Dough extends Product{
	Type: DoughType;
	Size: DoughSize;
	Prices: DoughPrices;
	constructor(options: Dougher = {}){
		super(options);

		this.Type = options.Type || DoughType.PAN;
		this.Size = options.Size || DoughSize.LARGE;
		this.Prices = options.Prices || {
											Small: {
												Thin: 5,
												Deep: 7,
												Pan: 6
											},
											Medium: {
												Thin: 7,
												Deep: 9,
												Pan: 8
											},
											Large: {
												Thin: 9,
												Deep: 11,
												Pan: 10
											}
										};
	}

	get Price(): number {
		return this.Prices[this.Size][this.Type];
	}

	set Price(price: number){
		console.log("not setting price");
	}

	static sizes(): string[]{
		return Object.keys(DoughSize);
	}
	static types(): string[]{
		return Object.keys(DoughType);
	}
	copy(): Dough{
		return new Dough(this);
	}
	getPrice(): number{
		return 
	}
}