import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';
import { User, UserAPI, UsersAPI } from '../models/user.model';
import { Error } from '../models/error.model';

@Injectable({
  providedIn: 'root',
})
export class UserService {

  private url: string = "http://localhost:8080/api";
  private http = inject(HttpClient);

  getUsers(): Observable<User[]> {
    return this.http.get<UsersAPI>(`${this.url}/users`)
      .pipe(map(
        (response: UsersAPI) => {
          return response.users;
        }
      ),
        catchError(
          (error) => {
            let err: Error = {
              status: error.error.status,
              error: error.error.error,
              details: error.error.details
            };
            return throwError(() => err);
          }
        )
      )
  }

  getUser(id: number): Observable<User> {
    return this.http.get<UserAPI>(`${this.url}/users/${id}`)
      .pipe(map(
        (response: UserAPI) => {
          return response.user;
        }
      ),
        catchError(
          (error) => {
            let err: Error = {
              status: error.error.status,
              error: error.error.error,
              details: error.error.details
            };
            return throwError(() => err);
          }
        )
      )
  }

  // TODO create user (expose backend)

}
