import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, catchError, map, Observable, of, tap } from 'rxjs';
import { AuthApi, Claims, Login } from '../models/auth.model';
import { Error } from '../models/error.model';

@Injectable({
  providedIn: 'root',
})
export class Auth {

  private url: string = "http://localhost:8080/api/auth";
  private http = inject(HttpClient)

  private logged = new BehaviorSubject<boolean>(this.hasToken())
  isLogged = this.logged.asObservable();

  getToken() : string | null {
    return localStorage.getItem("token");
  }

  hasToken() : boolean {
    return !!this.getToken();
  }

  setToken(token: string) : void {
    localStorage.setItem("token", token);
  }

  logout() : void {
    localStorage.removeItem("token");
  }

  getClaims() : Observable<Claims> {
    return this.http.get<AuthApi>(`${this.url}/verify`)
      .pipe(map(
        (response : AuthApi) => new Claims(response.claims)
      ));
  }

  testToken() : Observable<AuthApi | Error> {
    return this.http.get<AuthApi>(`${this.url}/verify`)
      .pipe(map(
        (response : AuthApi) => {
          return response
        }
      ),
      catchError(
        (error) => {
          let err : Error = {
            status: error.error.status,
            error: error.error.error,
            details: error.error.details
          };
          return of(err);
      })
    )
  }

  login(cred: {email:string, password:string}) : Observable<Login | Error> {
    return this.http.post<Login>(`${this.url}/login`, cred, {responseType:"json"})
      .pipe(map(
        (response : Login) => {
          this.setToken(response.token);
          return response;
        }
      ),
      catchError(
        (error) => {
          let err : Error = {
            status: error.error.status,
            error: error.error.error,
            details: error.error.details
          };
          return of(err);
        }
      )
    )
  }

}
