import { Inject, Component, OnInit, ViewChild, OnDestroy } from '@angular/core';
import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatAutocompleteTrigger } from '@angular/material';
import { Customer } from './customer.class';
import { FormControl } from '@angular/forms';

import { debounceTime, distinctUntilChanged, map, startWith } from 'rxjs/operators';
import { OrderService } from '../order.service';
import { User } from '../user/user.class';
import { Model } from '../common/model.class';
import { UserService } from '../user/user.service';
import { Subscription } from 'rxjs';
import { Router, ActivatedRoute } from '@angular/router';
import { CustomerLookupComponent } from './customer-lookup-dialog.component';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.component.html',
  styleUrls: ['./customer.component.css']
})
export class CustomerComponent {

    constructor(
        public dialog: MatDialog,
        private os: OrderService,
        private us: UserService,
        private router: Router,
        private route: ActivatedRoute
    ) { }

    search(e) {
        console.log(e);
        this.us.getCustomers().subscribe(list => {
            console.log(list);
            this.openDialog(list);
        });
    }

    openDialog(list: User[]) {
        const dialogRef = this.dialog.open(CustomerLookupComponent, {
                width: '250px',
                data: {
                        list: list
                }
    });

        dialogRef.afterClosed().subscribe(result => {
            if (result) {
                this.os.setCustomer(new User(result));
                console.log(result);
            }
                this.router.navigate(['../menu'], {relativeTo: this.route});
        });
    }
}
