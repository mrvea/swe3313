import { Model, Modelor } from 'src/app/common/model.class';

export enum Product_Category {// add to database
    DOUGH = 'Dough',
    SAUCE = 'Sauce',
    TOPPING = 'Topping',
    DRINK = 'Drink'
}

export interface Productor extends Modelor {
    Name?: string;
    Price?: number;
    Category?: string;
    Tooltip?: string;
    Image?: string;
}

export class Product extends Model implements Productor {
    Name: string;
    Price: number;
    Category: string;
    Tooltip?: string;
    Image: string;
    constructor(options: Productor = {}) {
        super(options);
        this.ID = options.ID || null;
        this.Name = options.Name || null;
        this.Category = options.Category || null;
        this.Price = options.Price || 0;
        this.Tooltip = options.Tooltip || null;
        this.Image = options.Image || 'default_pizza_image.png';
    }

    static getPrice(product): number {
        let price = 0;
        // console.log("product: ", product);
        if (product.constructor.name ===  'Pizza') {
            for (const key of ['Dough', 'Sauce', 'Toppings']) {
                // console.log(key);
                if (!product.hasOwnProperty(key)) {
                    continue;
                }
                const sec = product[key];
                // console.log(sec);
                if (key === 'Toppings') {
                    for (const topping of sec) {
                        price += topping.Price;
                    }

                } else {
                    price += sec.Price || 0;
                }

            }
            return price;
        }
        if (product instanceof Product) {
            return product.Price;
        }

        return price;
    }

    static getTotal(o: any[]): number {
        console.log(o);
        return o.map(p => Product.getPrice(p)).reduce((a, b) => a + b, 0);
    }
    copy(): Product {
        return new Product(this);
    }
}
