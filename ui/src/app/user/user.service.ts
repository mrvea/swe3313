import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from './user.class';

import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(
    protected http: HttpClient
  ) { }

  getCustomers(): Observable<User[]> {
    return this.http.get<User[]>('./CUSTOMER_DATA.json').pipe(map(this.toUser));
    // return this.http.get<User[]>('/api/customers');
  }

  toUser(list) {
    const temp = list.map(u => {
        return new User(u);
    });
    return temp;
}

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>('/api/users');
  }

  save(u: User): Observable<User> {
    if (!Boolean(u.ID)) {
        return this.update(u);
    }
    return this.http.post<User>('/api/user/save', u);
  }

  update(u: User): Observable<User> {
    return this.http.put<User>('/api/user/edit/' + u.ID, u);
  }

  delete(u: User): Observable<Boolean> {
    return this.http.delete<Boolean>('/api/user/delete/' + u.ID);
  }
}
