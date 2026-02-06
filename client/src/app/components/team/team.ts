import { Component, inject } from '@angular/core';
import { TeamService } from '../../services/team';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { AuthService } from '../../services/auth';
import { AuthApi } from '../../models/auth.model';
import { combineLatest, filter, map, startWith, switchMap } from 'rxjs';
import { AsyncPipe, NgClass, UpperCasePipe } from '@angular/common';

@Component({
  selector: 'app-team',
  imports: [AsyncPipe, ReactiveFormsModule, UpperCasePipe, NgClass],
  templateUrl: './team.html',
  styleUrl: './team.css',
})
export class Team {

  private teamSvc = inject(TeamService);
  private auth = inject(AuthService)
  searchCtrl = new FormControl('');

  uid$ = this.auth.getStatus().pipe(
    filter(Boolean),
    switchMap(() => this.auth.testToken()),
    map((resp: AuthApi) => resp.claims.userID)
  );

  teams$ = this.uid$.pipe(
    switchMap(uid => this.teamSvc.getOwnTeams(uid))
  );

  filteredTeams$ = combineLatest([
    this.teams$,
    this.searchCtrl.valueChanges.pipe(startWith(''))
  ]).pipe(
    map(([members, search]) => {
      const term = search?.toLowerCase().trim() ?? '';
      return members.filter(m =>
        m.team_name.toLowerCase().includes(term) ||
        m.role.toLowerCase().includes(term)
      );
    })
  );

}
