import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, catchError, map, Observable, of, throwError } from 'rxjs';
import { AuthApi, Claims, Login } from '../models/auth.model';
import { Error } from '../models/error.model';

@Injectable({
  providedIn: 'root',
})
export class AuthService {

  private url: string = "http://localhost:8080/api/auth";
  private http = inject(HttpClient)

  private logged = new BehaviorSubject<boolean>(this.hasToken())

  claims! : Claims;
  
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
    this.logged.next(false);
    localStorage.removeItem("token");
  }

  getStatus() : Observable<boolean> {
    return this.logged.asObservable();
  }

  getClaims() : Observable<Claims> {
    let current = new BehaviorSubject<Claims>(this.claims);
    return current.asObservable();
  }
  
  testToken() : Observable<AuthApi> {
    return this.http.get<AuthApi>(`${this.url}/verify`)
      .pipe(map(
        (response : AuthApi) => {
          this.claims = response.claims;
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
          return throwError(() => err);
      })
    )
  }

  login(cred: {email:string, password:string}) : Observable<Login> {
    return this.http.post<Login>(`${this.url}/login`, cred, {responseType:"json"})
      .pipe(map(
        (response : Login) => {
          this.setToken(response.token);
          this.logged.next(true);
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
          return throwError(() => err);
        }
      )
    )
  }

}
