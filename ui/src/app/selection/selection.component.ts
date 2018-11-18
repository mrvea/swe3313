import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Pizza, Dough, Topping, DoughType, DoughSize, ToppingType, SauceType, Product } from '../product/models';
import { OrderService } from '../order.service';
import { filter } from 'rxjs/operators';

const temp_pizzas: Product[] = [
    new Pizza({
        ID: 1,
        Name: 'Cheese Pizza',
        Price: 15.99,
        Dough: {
            ID: 11,
            Name: 'Thin Crush',
            Type:  DoughType.THIN,
        },
        Sauce: {
            ID: 12,
            Name: 'Marinara',
            Type: SauceType.RED,
            Price: 2
        },
        Toppings: [
            {
                ID: 13,
                Name: 'House Blend Cheese',
                Tooltip: 'mozzarella, provolone, cheddar and Parmesan',
                Type: ToppingType.CHEESE
            }
        ]
    }),
    new Pizza({
        ID: 2,
        Name: 'Pepperoni Pizza',
        Price: 17.99,
        Image: 'pizza_pepperoni.png',
        Dough: {
            ID: 11,
            Name: 'Pan Crust',
            Type:  DoughType.PAN,
        },
        Sauce: {
            ID: 12,
            Name: 'Marinara',
            Type: SauceType.RED,
            Price: 1.99
        },
        Toppings: [
            {
                ID: 23,
                Name: 'House Blend Cheese',
                Tooltip: 'mozzarella, provolone, cheddar and Parmesan',
                Type: ToppingType.CHEESE,
                Price: 1.99
            },
            {
                ID: 24,
                Name: 'Pepperoni',
                Type: ToppingType.MEAT,
                Price: 1.99
            }
        ]
    }),
    new Pizza({
        ID: 3,
        Name: 'Meat Lovers Pizza',
        Price: 19.99,
        Image: 'pan_meat_lovers.png',
        Dough: {
            ID: 11,
            Name: 'Pan Crust',
            Type:  DoughType.PAN,
        },
        Sauce: {
            ID: 12,
            Name: 'Marinara',
            Type: SauceType.RED,
            Price: 2
        },
        Toppings: [
            {
                ID: 13,
                Name: 'House Blend Cheese',
                Tooltip: 'mozzarella, provolone, cheddar and Parmesan',
                Type: ToppingType.CHEESE
            },
            {
                ID: 24,
                Name: 'Pepperoni',
                Type: ToppingType.MEAT
            },
            {
                ID: 35,
                Name: 'Sausage',
                Type: ToppingType.MEAT
            },
            {
                ID: 36,
                Name: 'Ham',
                Type: ToppingType.MEAT
            },
            {
                ID: 37,
                Name: 'Bacon',
                Type: ToppingType.MEAT
            }
        ]
    }),
    new Product({
        ID: 10,
        Name: 'Pepsi Cola',
        Price: 2.59
    })
];

const CRUSTS: Product[] = [
    new Dough({
        ID: 110,
        Name: 'Pan',
    }),
    new Dough({
        ID: 111,
        Name: 'Deep',
    }),
    new Dough({
        ID: 112,
        Name: 'Thin',
    })

];

const TOPPINGS = {
    Meat: [
        new Topping({
            ID: 1110,
            Name: 'Pepperoni'
        }),
        new Topping({
            ID: 1111,
            Name: 'Beef'
        }),
        new Topping({
            ID: 1112,
            Name: 'Sausage'
        }),
        new Topping({
            ID: 11121,
            Name: 'Ham'
        })

    ],
    Cheese: [
        new Topping({
            ID: 1113,
            Name: 'House'
        }),
        new Topping({
            ID: 1114,
            Name: 'Mozzarella'
        }),
        new Topping({
            ID: 1115,
            Name: 'Provolone'
        }),
        new Topping({
            ID: 11151,
            Name: 'Parmesan'
        })
        ,
        new Topping({
            ID: 11152,
            Name: 'Cheddar'
        })
    ],
    Vegie: [
        new Topping({
            ID: 1116,
            Name: 'Tomato'
        }),
        new Topping({
            ID: 1117,
            Name: 'Spinach'
        }),
        new Topping({
            ID: 1118,
            Name: 'Pepper'
        }),
        new Topping({
            ID: 1118,
            Name: 'Onion'
        })
    ]
};

@Component({
  selector: 'app-selection',
  templateUrl: './selection.component.html',
  styleUrls: ['./selection.component.css']
})
export class SelectionComponent implements OnInit {
    private _pIndex = 0;
    pizzas: Pizza[];
    products: any;
    stage: any;
    doughSize = DoughSize;
    private _stages = [
        {
            stage: 'Menu',
            title: 'Please suggest pre-made pizza. Or make a custom pizza.',
            method: 'getMenu'
        },
        {
            stage: 'Crust',
            title: 'Please ask about any updates to the crust and/or sause.',
            method: 'getCrustSizeSause'
        },
        {
            stage: 'Toppings',
            title: 'Please ask about any additional toppings.',
            method: 'getToppings'
        }
    ];
  constructor(
      protected os: OrderService,
      private route: ActivatedRoute
  ) {
      this.route.data.subscribe((data: {
          stage: string,
          title: string
      }) => {
          if (!Boolean(data.stage)) {
              this.setStage('Menu');
              return;
          }
          this.setStage(data.stage);
      });
  }

  ngOnInit() {
      // this.setStage("menu");
      this.route.queryParams.pipe(filter(params => params.p)).subscribe(params => {
          console.log(params);
          this._pIndex = params.p;
      });
  }

  setStage(stage) {
      this.stage = this._stages.find(st => st.stage === stage);
      this[this.stage.method]();
  }

  getMenu() {
      console.log('Setting Menu');
      this.products = temp_pizzas;
  }

  getCrustSizeSause() {
      this.products = CRUSTS;
  }

  getToppings() {
      this.products = TOPPINGS;
  }
  add(product) {
      switch (this.stage.stage) {
          case 'Toppings':
              console.log('addng to toppings: ', product);
              this.os.addToppingTo(this._pIndex, product.copy());
              break;

          default:
              this.os.addToOrder(product.copy());
              break;
      }

  }

  getDoughSize(): string[] {
      return Dough.sizes();
  }

  getPrice(product): number {
      return Product.getPrice(product);
  }

  getKeys(item: {}) {
      return Object.keys(item);
  }

  isArray(item: any) {
      return Array.isArray(item);
  }
}
