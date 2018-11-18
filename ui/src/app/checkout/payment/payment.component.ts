import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { FormBuilder, FormGroup, Validators, FormControl } from '@angular/forms';
import { map, distinctUntilChanged, debounceTime, tap, delay } from 'rxjs/operators';
import { CheckoutService } from '../checkout.service';
import { Router } from '@angular/router';
import { OrderService, SendPayload } from 'src/app/order.service';
import { Order } from 'src/app/order.class';
import { Product } from 'src/app/product/models';
import { User } from 'src/app/user/user.class';

const cardTypeMap = {
  Visa: {
    code: 3,
    start: [4]
  },
  MasterCard: {
    code: 3,
    start: [51, 52, 53, 54, 55]
  },
  Discovery: {
    code: 3,
    start: [6011, 65]
  },
  Amex: {
    code: 4,
    start: [34, 37]
  }
};

const cardStarts = {
  4: 'Visa',
  51: 'MasterCard',
  52: 'MasterCard',
  53: 'MasterCard',
  54: 'MasterCard',
  55: 'MasterCard',
  65: 'Discover',
  6011: 'Discover',
  34: 'Amex',
  37: 'Amex',
};

@Component({
  selector: 'app-payment',
  templateUrl: './payment.component.html',
  styleUrls: ['./payment.component.css']
})
export class PaymentComponent implements OnInit {
  form: FormGroup;
  givenValue: FormControl = new FormControl();
  order: Order = new Order();
  card: any = {};
  isManual = false;
  isCash = false;
  fields = {
    Name: {
      Name: 'Name',
      Value: ['', Validators.required],
      Label: 'Name'
    },
    Number: {
      Name: 'Number',
      Label: 'Card Number',
      Value: ['', [Validators.required, Validators.maxLength(16)]]
    },
    Exp: {
      Name: 'Exp',
      Label: 'Expiration Date',
      Value: ['', Validators.required]
    },
    Code: {
      Name: 'Code',
      Label: 'Security Code',
      Value: ['', Validators.required]
    }
  };

  constructor(
    private fb: FormBuilder,
    private location: Location,
    private cs: CheckoutService,
    private router: Router,
    private os: OrderService
  ) {
    const group = {};
    // tslint:disable-next-line:forin
    for (const key in this.fields) {
        const field = this.fields[key];
        group[field.Name] = field.Value;
    }
    this.form = this.fb.group(group);
  }

  ngOnInit(): void  {
    this.form.controls.Number.valueChanges.pipe(
      debounceTime(500),
      distinctUntilChanged(),
      tap(val => console.log(val)),
      map(val => {
      // tslint:disable-next-line:forin
        for (const key in cardStarts) {
          const card = cardStarts[key];
          if (val.indexOf(key) === 0) {
          return Object.assign({}, cardTypeMap[card], {name: card});
          }
        }
      })
    ).subscribe(val => {
      this.card = val;
      this.order.CardType = this.card.name;
    });

    this.os.order$.pipe(
      tap((p: SendPayload) => this.order.TotalPrice = Product.getTotal(p.product))
    ).subscribe(p => {
      this.order.Products = p.product;
    });

    this.os.customer$.subscribe( c => {
      this.order.Customer = c;
    });
  }

  processCard(e) {
    this.order.Type = 'card';
    this.cs.processCard(null).pipe(
      tap(_ => alert('Card machine is ready')),
      delay(2000),
      tap(_ => alert('Authorizing...')),
      delay(1000),
      tap(_ => alert('Approved...')),
      delay(500),
      tap(_ => alert('Printing Reciept...'))
    ).subscribe(resp => {
      this.os.reset();
      this.router.navigate(['dash']);
    });
  }

  processCheck(e) {
    this.order.Type = 'check';
    this.cs.processCard(this.order).pipe(
      tap(_ => alert('Card machine is ready')),
      delay(2000),
      tap(_ => alert('Authorizing...')),
      delay(1000),
      tap(_ => alert('Approved...')),
      delay(500),
      tap(_ => alert('Printing Reciept...'))
    ).subscribe(resp => {
      this.os.reset();
      this.router.navigate(['dash']);
    });
  }

  setManual(e) {
    this.isManual = true;
  }

  processCash(e) {
    this.isCash = true;
    this.order.Type = 'cash';
  }

  getChange() {
    console.log(this.order);
    console.log(this.givenValue.value);
    this.givenValue.patchValue(`Change ${+this.givenValue.value - this.order.TotalPrice}`);
  }

  goBack() {
    this.location.back();
  }

  done() {
    this.os.save(this.order);
    this.os.reset();
    this.router.navigate(['dash']);
  }

  onSubmit(e): void {
    if (this.form.invalid) {
      alert('Please check the form and make sure everything is filled in correctly');
      return;
    }
    this.order.Customer = new User(this.form.value);
    this.order.Type = 'card-manual';
    this.cs.processCard(this.order).pipe(
      tap(_ => alert('Authorizing...')),
      delay(1000),
      tap(_ => alert('Approved...')),
      delay(500),
      tap(_ => alert('Printing Reciept...'))
    ).subscribe(resp => {
      this.os.reset();
      this.router.navigate(['dash']);
    });
  }
}
