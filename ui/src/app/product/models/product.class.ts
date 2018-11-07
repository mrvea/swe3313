export enum Product_Category {// add to database
	DOUGH = "Dough",
	SAUCE = "Sauce",
	TOPPING = "Topping",
	DRINK = "Drink"
}

export interface Productor{
	ID?: number,
	Name?: string,
	Price?: number,
	Category?: string,
	Created?: string | Date,
	Modified?: string | Date,
	Tooltip?: string,
	Image?: string
}

export class Product implements Productor{
	ID: number;
	Name: string;
	Price: number;
	Category: string;
	Created?: Date;
	Modified?: Date;
	Tooltip?: string;
	Image: string;
	constructor(options: Productor = {}){
		this.ID = options.ID || null;
		this.Name = options.Name || null;
		this.Category = options.Category || null;
		this.Price = options.Price || 0;
		this.Tooltip = options.Tooltip || null;
		this.Image = options.Image || "default_pizza_image.png";
		this.Created = this.isDate(options.Created)? 
						options.Created : this.makeDate(options.Created);
		this.Modified = this.isDate(options.Modified)? 
						options.Modified : this.makeDate(options.Modified);
	}

	isDate(d): d is Date{
		if(d instanceof Date){
			return true;
		}
		return false;
	}

	makeDate(strDate: string | null): Date{
		if(strDate == null){
			return new Date();
		}
		return new Date(strDate);
	}
}