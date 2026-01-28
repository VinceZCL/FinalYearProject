import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';
import { User, UserAPI, UsersAPI } from '../models/user.model';
import { Error } from '../models/error.model';

@Injectable({
  providedIn: 'root',
})
export class UserService {

  private url: string = "http://localhost:8080/api/users";
  private http = inject(HttpClient);

  getUsers(): Observable<User[]> {
    return this.http.get<UsersAPI>(`${this.url}`)
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
    return this.http.get<UserAPI>(`${this.url}/${id}`)
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

  newUser(cred: { name: string, email: string, password: string, type: string }): Observable<UserAPI> {
    return this.http.post<UserAPI>(`${this.url}`, cred, { responseType: "json" })
      .pipe(map(
        (resp: UserAPI) => {
          return resp;
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
      );
  }

  deactivateUser(id: number): Observable<User> {
    return this.http.delete<UserAPI>(`${this.url}/${id}`)
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

}
