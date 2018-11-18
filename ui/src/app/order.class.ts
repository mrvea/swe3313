import { Model } from './common/model.class';
import { Product } from './product/models';
import { User } from './user/user.class';

type OrderStatus = 'waiting' | 'processing' | 'submitting' | 'confirming';

export class Order extends Model {
    Products: any[];
    CardType: string;
    Type: string;
    Status: OrderStatus;
    TotalPrice: number;
    Customer: User;
    constructor(options: {
        ID?: number,
        Products?: any[],
        CardType?: string,
        Type?: string,
        Created?: string | Date,
        Modified?: string | Date,
        Status?: OrderStatus,
        TotalPrice?: number,
        Customer?: User
    } = {Products: []}) {
        super(options);

        this.Products = options.Products || [];
        this.CardType = options.CardType || '';
        this.Type = options.Type || '';
        this.Status = options.Status || 'waiting';
        this.TotalPrice = options.TotalPrice || 0;
        this.Customer = options.Customer || null;
    }
}
