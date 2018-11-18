import { Injectable } from '@angular/core';
import { Product } from './models';
import { Observable, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  constructor(
        protected http: HttpClient
  ) { }

  getMenu(): Observable<Product[]> {
    return of([]);
  }

  getCrusts(): Observable<Product[]> {
    return of([]);
  }

  getToppings(): Observable<Product[]> {
    return of([]);
  }

  save(p: Product): Observable<Product> {
    if (!Boolean(p.ID)) {
        return this.update(p);
    }
    return this.http.post<Product>('/api/product/save', p);
  }

  update(p: Product): Observable<Product> {
    return this.http.put<Product>('/api/product/edit/' + p.ID, p);
  }

  delete(p: Product): Observable<Boolean> {
    return this.http.delete<Boolean>('/api/product/delete/' + p.ID);
  }
}
