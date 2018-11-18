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
  MatSidenavModule, MatDialogModule, MatInputModule,
  MatDividerModule, MatDatepickerModule, MatNativeDateModule, MatAutocompleteModule
} from '@angular/material';
import { UserComponent } from './user/user.component';
import { CustomerComponent } from './customer/customer.component';
import { ProductComponent } from './product/product.component';
import { CheckoutComponent } from './checkout/checkout.component';
import { PaymentComponent } from './checkout/payment/payment.component';
import { DeliveryComponent } from './checkout/delivery/delivery.component';
import { PageNotFoundComponent } from './common/page-not-found/page-not-found.component';
import { CustomerLookupComponent } from './customer/customer-lookup-dialog.component';

const APP_ROUTES: Routes = [
  {
    path: '',
    redirectTo: 'dash',
    pathMatch: 'full'
  },
  {
    path: 'dash',
    children: [
      {
        path: '',
        redirectTo: 'customer',
        pathMatch: 'full'
      },
      {
        path: 'customer',
        component: CustomerComponent
      },
      {
        path: 'menu',
        component: SelectionComponent,
        data: {
          state: 'Menu',
          title: 'Menu'
        }
      },
      {
        path: 'crust',
        component: SelectionComponent,
        data: {
          stage: 'Crust',
          title: 'Please choose a crust',
        }
      },
      {
        path: 'toppings',
        component: SelectionComponent,
        data: {
          stage: 'Toppings',
          title: 'Please select a topping'
        }
      },
      {
        path: 'checkout',
        children: [
          {
            path: '',
            component: CheckoutComponent,
          },
          {
            path: 'payment',
            component: PaymentComponent
          },
          {
            path: 'delivery',
            component: DeliveryComponent
          }
        ]
      },
      {
          path: '**',
          component: PageNotFoundComponent
      }
    ]
  }
];

@NgModule({
  declarations: [
    AppComponent,
    SelectionComponent,
    DetailsComponent,
    UserComponent,
    CustomerComponent,
    ProductComponent,
    CheckoutComponent,
    CustomerLookupComponent,
    PaymentComponent,
    DeliveryComponent,
    PageNotFoundComponent
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
    MatInputModule,
    MatDividerModule,
    MatDatepickerModule,
    MatNativeDateModule,
    HttpClientModule,
    MatAutocompleteModule
  ],
  providers: [],
  entryComponents: [CustomerLookupComponent],
  bootstrap: [AppComponent]
})
export class AppModule { }
