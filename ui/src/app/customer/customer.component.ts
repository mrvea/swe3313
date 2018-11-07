import { Inject, Component, OnInit } from '@angular/core';
import { MatDialog, MAT_DIALOG_DATA, MatDialogRef } from '@angular/material';
import { Customer } from './customer.class';
import { FormControl } from '@angular/forms';

import { debounceTime, distinctUntilChanged, map, startWith } from 'rxjs/operators';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.component.html',
  styleUrls: ['./customer.component.css']
})
export class CustomerComponent implements OnInit {

  constructor(
  	public dialog: MatDialog
  ) { }

  ngOnInit() {
  }

  search(e){
  	console.log(e);
  	const dialogRef = this.dialog.open(CustomerLookupComponent, {
      width: '250px',
      data: new Customer()
    });

    dialogRef.afterClosed().subscribe(result => {
      if(result) {
          console.log(result);
      }
    });
  }
}

const format = {
	pattern: [3, 3, 4],
	open: ["(", ") ", "-"]
}

@Component({
  selector: 'customer-lookup-dialog',
  templateUrl: './customer-lookup-dialog.component.html',
})

export class CustomerLookupComponent implements OnInit {
   	searchValue: FormControl
    constructor(
      public dialogRef: MatDialogRef<CustomerLookupComponent>,
      @Inject(MAT_DIALOG_DATA) public data: Customer
    ) {
    	this.searchValue = new FormControl();
    }

    ngOnInit():void {
    	this.searchValue.valueChanges.pipe(
    		debounceTime(500),
    		map(val => {
    			val = val.replace(/\D/g, '').replace(/\s/g, "");
    			return val;
    		}),
    		distinctUntilChanged()
    	).subscribe(v => {
    		console.log(v);
    		this.searchValue.patchValue(this.formatPhone(v));
    	})
    }

    formatPhone(val, depth: number = 0){
    	console.log(val);
    	console.log(depth);
    	if(val.length == 0){
    		return "";
    	}
    	const n = val.substring(0, format.pattern[depth]);

    	return format.open[depth] + n + this.formatPhone(val.substring(format.pattern[depth]), ++depth);
    }

    onNoClick(): void {
      this.dialogRef.close();
    }

     done() {
    }

}
