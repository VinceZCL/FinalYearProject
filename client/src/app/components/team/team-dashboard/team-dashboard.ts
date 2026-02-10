import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { TeamService } from '../../../services/team';
import { ActivatedRoute, Router } from '@angular/router';
import { Member, Team } from '../../../models/team.model';
import { Error } from '../../../models/error.model';

@Component({
  selector: 'app-team-dashboard',
  imports: [],
  templateUrl: './team-dashboard.html',
  styleUrl: './team-dashboard.css',
})
export class TeamDashboard implements OnInit {

  private teamSvc = inject(TeamService);
  private route = inject(ActivatedRoute);
  private cd = inject(ChangeDetectorRef);
  private router = inject(Router);

  teamID!: number;
  team!: Team;
  members!: Member[];

  ngOnInit(): void {
    this.teamID = Number(this.route.snapshot.paramMap.get("teamID"));
    this.teamSvc.getTeam(this.teamID).subscribe({
      next: (resp: Team) => {
        this.team = resp;
        this.getMembers();
      },
      error: (err : Error) => {
        console.log(err.details);
        this.router.navigate(["/404"]);
      }
    })
  }

  // ! get own role
  // ! if role is ADMIN, able to add people
  // ? call member API to check own role

  getMembers(): void {
    this.teamSvc.getMembers(this.teamID).subscribe({
      next: (resp: Member[]) => {
        this.members = resp;
        this.cd.detectChanges();
      },
      error: (err: Error) => {
        console.log(err.details);
      }
    })
  }

}
