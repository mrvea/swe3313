import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { FormGroup, Validators, FormBuilder } from '@angular/forms';
import { debounceTime, map, distinctUntilChanged } from 'rxjs/operators';
import { Model } from 'src/app/common/model.class';
import { User } from 'src/app/user/user.class';
import { UserService } from 'src/app/user/user.service';
import { OrderService } from 'src/app/order.service';

@Component({
  selector: 'app-delivery',
  templateUrl: './delivery.component.html',
  styleUrls: ['./delivery.component.css']
})
export class DeliveryComponent implements OnInit {
  form: FormGroup;
  customer: User = new User();
  fields = {
    Name: {
      Name: 'FullName',
      Label: 'Full Name',
      Value: ['', Validators.required]
    },
    Email: {
      Name: 'Email',
      Label: 'Email',
      Value: ''
    },
    Address: {
      Name: 'Address',
      Label: 'Address',
      Value: ['', Validators.required]
    },
    Address2: {
      Name: 'Address2',
      Label: 'Address 2',
      Value: ''
    },
    City: {
      Name: 'City',
      Label: 'City',
      Value: ''
    },
    Phone: {
      Name: 'Phone',
      Label: 'Phone',
      Value: ['', Validators.required],
    }
  };

  constructor(
    private fb: FormBuilder,
    private us: UserService,
    private location: Location,
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

  ngOnInit() {
    this.form.controls[this.fields.Name.Name].valueChanges.pipe(
      map((val: string) => val.split(' '))
    ).subscribe(name => {
      console.log(name);
      this.customer.FirstName = name[0].trim();
      if (name.length === 1) {
        return;
      }
      this.customer.LastName = name[1].trim();
    });

    this.form.controls.Phone.valueChanges.pipe(
      debounceTime(500),
      map(val => Number.isInteger(val) ? val.toString() : val),
      map(val => {
          console.log(val);
          val = val.replace(/\D/g, '').replace(/\s/g, '');
          return val;
      }),
      distinctUntilChanged()
    ).subscribe(v => {
        console.log(v);
        this.form.controls.Phone.patchValue(Model.formatPhone(v));
    });

    this.os.customer$.subscribe(c => {
      this.customer = c;
      this.form.patchValue(c);
    });
  }

  onSubmit(e) {
    if (this.form.invalid) {
      alert('Please make sure the form fully filled!');
      return;
    }
    Object.assign(this.customer, this.form.value);
    this.us.save(this.customer).subscribe(resp => {
      if (resp) {
        alert('Save successful!');
      }
    });
  }

  goBack() {
    this.location.back();
  }

}
