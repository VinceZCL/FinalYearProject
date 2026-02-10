import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, catchError, map, Observable, of, tap, throwError } from 'rxjs';
import { AuthApi, Claims, Login } from '../models/auth.model';
import { Error } from '../models/error.model';

@Injectable({
  providedIn: 'root',
})
export class AuthService {

  private url: string = "http://localhost:8080/api/auth";
  private http = inject(HttpClient)

  private claimSubject = new BehaviorSubject<Claims | null>(null);
  claim$ = this.claimSubject.asObservable();

  private userIDSubject = new BehaviorSubject<Number | null>(null);
  userID$ = this.userIDSubject.asObservable();

  private loggedSubject = new BehaviorSubject<boolean>(false);
  logged$ = this.loggedSubject.asObservable();

  private adminSubject = new BehaviorSubject<boolean>(false);
  admin$ = this.adminSubject.asObservable();

  verifyAndHydrate(): Observable<boolean> {
    if (!this.hasToken()) {
      this.clearAuthState();
      return of(false);
    }

    return this.testToken().pipe(
      tap(
        (resp: AuthApi) => {
          this.claimSubject.next(resp.claims);
          this.userIDSubject.next(resp.claims.userID);
          this.loggedSubject.next(true);
          this.adminSubject.next(resp.claims.type == "admin");
        }
      ),
      map(() => true),
      catchError(() => {
        this.clearAuthState();
        return of(false);
      })
    );
  }

  clearAuthState(): void {
    this.claimSubject.next(null);
    this.userIDSubject.next(null);
    this.loggedSubject.next(false);
  }

  logout(): void {
    localStorage.removeItem("token");
    this.clearAuthState();
  }

  getToken(): string | null {
    return localStorage.getItem("token");
  }

  hasToken(): boolean {
    return !!this.getToken();
  }

  setToken(token: string): void {
    localStorage.setItem("token", token);
  }

  testToken(): Observable<AuthApi> {
    return this.http.get<AuthApi>(`${this.url}/verify`)
      .pipe(map(
        (response: AuthApi) => {
          return response
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
          })
      )
  }

  login(cred: { email: string, password: string }): Observable<Login> {
    return this.http.post<Login>(`${this.url}/login`, cred, { responseType: "json" })
      .pipe(map(
        (response: Login) => {
          this.setToken(response.token);
          this.loggedSubject.next(true);
          return response;
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
