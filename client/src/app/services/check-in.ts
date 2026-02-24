import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';
import { CheckIns, CheckInsAPI, NewCheckIns } from '../models/check-in.model';
import { Error } from '../models/error.model';

@Injectable({
  providedIn: 'root',
})
export class CheckInService {

  private url: string = "http://localhost:8080/api/checkins";
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
      )
  }

  submitBulk(payload: {"checkIns": NewCheckIns[]}): Observable<CheckIns | null> {
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
      )
  }

}
