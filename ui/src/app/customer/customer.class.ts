

export class Address {
	Street: string;
	City: string;
	Street2: string;

	constructor(options: {
		Street?: string,
		City?: string,
		Street2?: string
	} = {}){
		this.Street = options.Street || "";
		this.City = options.City || "";
		this.Street2 = options.Street2 || "";
	}
}
export class Customer{
	Name: string;
	Phone: string;
	CardType: string;
	Address: Address;
	constructor(options: {
		Name?: string,
		Phone?: string,
		CardType?: string,
		Address?: Address
	} = {}){
		this.Name = options.Name || "";
		this.Phone = options.Phone || "";
		this.CardType = options.Phone || "";
		this.Address = this.Address || new Address();
	}
}