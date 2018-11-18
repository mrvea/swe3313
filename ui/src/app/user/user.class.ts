import { Model } from '../common/model.class';

export class User extends Model {
    FirstName: string;
    LastName: string;
    FullName: string;
    Email: string;
    Role: string;
    Phone: number;
    Address: string;
    Address2: string;
    City: string;
    Zip: number;

    constructor(options: {
        ID?: number,
        FirstName?: string,
        LastName?: string,
        Role?: string,
        Phone?: number,
        Email?: string,
        Address?: string,
        Address2?: string,
        City?: string,
        Zip?: number
    }= {}) {
        super(options);
        this.FirstName = options.FirstName || '';
        this.LastName = options.LastName || '';
        this.FullName = `${this.FirstName} ${this.LastName}`;
        this.Role = options.Role || 'customer';
        this.Phone = +options.Phone || null;
        this.Email = options.Email || '';
        this.Address = options.Address || '';
        this.Address2 = options.Address2 || '';
        this.City = options.City || '';
    }
}
