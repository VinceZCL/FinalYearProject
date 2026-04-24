import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';
import { CheckIn, CheckIns, CheckInsAPI, NewCheckIns, TeamCheckInsAPI } from '../models/check-in.model';
import { Error } from '../models/error.model';
import { environment } from '../../environments/environments';

@Injectable({
  providedIn: 'root',
})
export class CheckInService {

  private url: string = `${environment.apiBase}/api/checkins`;
  private http = inject(HttpClient);

  // get uid from component via auth
  getDaily(uid: number): Observable<CheckIns | null> {
    return this.http.get<CheckInsAPI>(`${this.url}/users/${uid}`)
      .pipe(map(
        (resp: CheckInsAPI) => {
          return resp.checkIns;
          // null if no checkIn today
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

  submitBulk(payload: { "checkIns": NewCheckIns[] }): Observable<CheckIns | null> {
    return this.http.post<CheckInsAPI>(`${this.url}/bulk`, payload, { responseType: "json" })
      .pipe(map(
        (resp: CheckInsAPI) => {
          return resp.checkIns;
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

  getTeam(tid: number): Observable<CheckIns[]> {
    return this.http.get<TeamCheckInsAPI>(`${this.url}/teams/${tid}`)
      .pipe(map(
        (resp: TeamCheckInsAPI) => {
          return resp.checkIns;
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

  getYesterday(uid: number): Observable<CheckIn[] | undefined> {
    let d = new Date();
    d.setDate(d.getDate() - 1);
    const yesterday = d.toLocaleDateString('en-CA');
    return this.http.get<CheckInsAPI>(`${this.url}/users/${uid}?date=${yesterday}`)
      .pipe(map(
        (resp: CheckInsAPI) => {
          return resp.checkIns?.today;
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

  getDate(uid: number, date: string): Observable<CheckIns | null> {
    return this.http.get<CheckInsAPI>(`${this.url}/users/${uid}?date=${date}`)
      .pipe(map(
        (resp: CheckInsAPI) => {
          return resp.checkIns;
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
