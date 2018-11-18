import { User } from '../user/user.class';
import { FormControl } from '@angular/forms';
import { Subscription } from 'rxjs';
import { Component, OnInit, OnDestroy, Inject, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA, MatAutocompleteTrigger } from '@angular/material';
import { startWith, debounceTime, map, distinctUntilChanged } from 'rxjs/operators';
import { Model } from '../common/model.class';

@Component({
    // tslint:disable-next-line:component-selector
    selector: 'customer-lookup-dialog',
    templateUrl: './customer-lookup-dialog.component.html',
  })

export class CustomerLookupComponent implements OnInit, OnDestroy {
    searchValue: FormControl;
    list: User[];
    subscription: Subscription;
    constructor(
        public dialogRef: MatDialogRef<CustomerLookupComponent>,
        @Inject(MAT_DIALOG_DATA) public data: { list: User[] }
    ) {
        this.searchValue = new FormControl();
        this.list = [];
    }

    @ViewChild('autoCompleteInput', { read: MatAutocompleteTrigger}) autoComplete: MatAutocompleteTrigger;

    ngOnInit(): void {
        console.log(Object.assign({}, this.data));
        this.subscription = this.searchValue.valueChanges.pipe(
            startWith(''),
            debounceTime(500),
            map(val => {
                val = val.replace(/\D/g, '').replace(/\s/g, '');
                return val;
            }),
            distinctUntilChanged()
        ).subscribe(v => {
            console.log(v);
            this.list = this.data.list.filter(item => item.Phone.toString().indexOf(v) === 0);
            this.searchValue.patchValue(Model.formatPhone(v));
        });
    }

    ngOnDestroy(): void {
        this.subscription.unsubscribe();
    }

    onNoClick(): void {
    this.dialogRef.close();
    }

    displayFn(value) {
        if (!value || !value.Phone) {
            console.log(value);
            return value;
        }
            return Model.formatPhone(value.Phone.toString());
    }

    toggle(e, obj) {
        e.preventDefault();
        e.stopPropagation();
        console.log(e);
        console.log(obj);
        if (obj.isOpen) {
            this.autoComplete.closePanel();
            return;
        }
        this.autoComplete.openPanel();
    }
    select(e) {
        this.subscription.unsubscribe();
                    this.dialogRef.close(e.option.value);
    }

    done() {
    }

  }
