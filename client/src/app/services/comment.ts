import { inject, Injectable } from '@angular/core';
import { environment } from '../../environments/environments';
import { HttpClient } from '@angular/common/http';
import { NewComment, NewCommentAPI } from '../models/comment';
import { Error } from '../models/error.model';
import { catchError, map, Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class CommentService {

  private url: string = `${environment.apiBase}/api/comments`;
  private http = inject(HttpClient);

  newCommment(payload: NewComment): Observable<void> {
    return this.http.post<NewCommentAPI>(`${this.url}`, payload, { responseType: "json" })
      .pipe(map(
        (resp: NewCommentAPI) => {
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

}
