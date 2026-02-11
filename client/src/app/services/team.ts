import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';
import { Member, MemberAPI, MembersAPI, Team, TeamAPI, TeamCreatedAPI } from '../models/team.model';
import { Error } from '../models/error.model';

@Injectable({
  providedIn: 'root',
})
export class TeamService {
  
  private url: string = "http://localhost:8080/api/teams";
  private memUrl: string = "http://localhost:8080/api/members";
  private http = inject(HttpClient);

  getOwnTeams(userID: number): Observable<Member[]> {
    return this.http.get<MembersAPI>(`${this.url}/users/${userID}`)
      .pipe(map(
        (response: MembersAPI) => {
          return response.members;
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

  newTeam(cred: {name: string, creatorID: number}): Observable<TeamCreatedAPI> {
    return this.http.post<TeamCreatedAPI>(`${this.url}`, cred, {responseType: "json"})
      .pipe(map(
        (resp: TeamCreatedAPI) => {
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

  getTeam(teamID: number): Observable<Team> {
    return this.http.get<TeamAPI>(`${this.url}/${teamID}`)
      .pipe(map(
        (response: TeamAPI) => {
          return response.team;
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

  getMembers(teamID: number): Observable<Member[]> {
    return this.http.get<MembersAPI>(`${this.url}/members/${teamID}`)
      .pipe(map(
        (response: MembersAPI) => {
          return response.members;
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

  newMember(cred: {userID: number, teamID: number, role: string}): Observable<MemberAPI> {
    return this.http.post<MemberAPI>(`${this.memUrl}`, cred, { responseType: "json" })
      .pipe(map(
        (response: MemberAPI) => {
          return response;
        }
      ),
      catchError(
        (error) => {
          let err: Error = {
            status: error.error.status,
            error: error.error.error,
            details: error.error.details,
          };
          return throwError(() => err);
        }
      )
    )
  }

}
