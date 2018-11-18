import { Product, Productor } from './product.class';

export interface Toppinger extends Productor {
    Type?: ToppingType;
}

export enum ToppingType {// database?
    MEAT = 'Meat',
    VEG = 'Vegetable',
    CHEESE = 'Cheese',
    UNKNOWN = 'Unknown'
}
export class Topping extends Product {
    Type: ToppingType;
    constructor(options: Toppinger = {}) {
        super(options);

        this.Type = options.Type || ToppingType.UNKNOWN;
        this.Price = .99;
    }
    copy(): Topping {
        return new Topping(this);
    }
}
