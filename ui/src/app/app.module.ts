import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule, Routes } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { SelectionComponent } from './selection/selection.component';
import { DetailsComponent } from './details/details.component';

import {
  MatToolbarModule, MatFormFieldModule, MatSelectModule,
  MatGridListModule, MatIconModule, MatButtonModule,
  MatMenuModule, MatListModule, MatCardModule,
  MatSidenavModule, MatDialogModule, MatInputModule
} from '@angular/material';
import { UserComponent } from './user/user.component';
import { CustomerComponent, CustomerLookupComponent } from './customer/customer.component';
import { ProductComponent } from './product/product.component';
import { CheckoutComponent } from './checkout/checkout.component';

const APP_ROUTES: Routes = [
  {
    path: "selection",
    component: SelectionComponent
  },
  {
    path: "",
    component: CustomerComponent
  },
  {
    path: "checkout",
    component: CheckoutComponent
  }
]

@NgModule({
  declarations: [
    AppComponent,
    SelectionComponent,
    DetailsComponent,
    UserComponent,
    CustomerComponent,
    ProductComponent,
    CheckoutComponent, 
    CustomerLookupComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    RouterModule.forRoot(APP_ROUTES),
    ReactiveFormsModule,
    MatToolbarModule,
    MatFormFieldModule,
    MatSelectModule,
    MatGridListModule,
    MatIconModule,
    MatButtonModule,
    MatMenuModule,
    MatListModule,
    MatCardModule,
    MatSidenavModule,
    MatDialogModule,
    MatInputModule
  ],
  providers: [],
  entryComponents: [CustomerLookupComponent],
  bootstrap: [AppComponent]
})
export class AppModule { }
