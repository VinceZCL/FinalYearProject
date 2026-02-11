import { ChangeDetectorRef, Component, HostListener, inject, OnInit } from '@angular/core';
import { TeamService } from '../../../services/team';
import { ActivatedRoute, Router } from '@angular/router';
import { Member, MemberAPI, Team } from '../../../models/team.model';
import { Error } from '../../../models/error.model';
import { AuthService } from '../../../services/auth';
import { User } from '../../../models/user.model';
import { UserService } from '../../../services/user';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-team-dashboard',
  imports: [FormsModule],
  templateUrl: './team-dashboard.html',
  styleUrl: './team-dashboard.css',
})
export class TeamDashboard implements OnInit {

  private teamSvc = inject(TeamService);
  private auth = inject(AuthService);
  private userSvc = inject(UserService);
  private route = inject(ActivatedRoute);
  private cd = inject(ChangeDetectorRef);
  private router = inject(Router);

  uid!: number;
  teamID!: number;
  team!: Team;
  members!: Member[];
  power!: boolean;
  avail!: User[];

  // UI state
  sidebarOpen = false;
  search = '';
  filteredAvail: User[] = [];

  selectedUser?: User;
  showUserModal = false;

  ngOnInit(): void {
    this.teamID = Number(this.route.snapshot.paramMap.get("teamID"));
    this.auth.claim$.subscribe(claims => {
      this.uid = Number(claims?.userID);
    })
    this.teamSvc.getTeam(this.teamID).subscribe({
      next: (resp: Team) => {
        this.team = resp;
        this.getMembers();
      },
      error: (err : Error) => {
        console.log(err.details);
        this.router.navigate(["/404"]);
      }
    });
  }

  getMembers(): void {
    this.teamSvc.getMembers(this.teamID).subscribe({
      next: (resp: Member[]) => {
        this.members = resp;
        this.power = this.members.some(
          (m: Member) => {
            return m.userID === this.uid && m.role === "admin";
          }
        );
        this.getAvailable();
      },
      error: (err: Error) => {
        console.log(err.details);
      }
    });
  }

  getAvailable(): void {
    let memberIDs = new Set(this.members.map(m => m.userID));
    this.userSvc.getUsers().subscribe({
      next: (val: User[]) => {
        this.avail = val.filter(
          u => !memberIDs.has(u.id) && u.status !== "deactivated"
        );
        this.cd.detectChanges();
      }
    });
  }

  onSearchChange(): void {
    const term = this.search.trim().toLowerCase();
    if (!term) {
      this.filteredAvail = [];
      return;
    }

    this.filteredAvail = this.avail.filter(u =>
      `${u.name} ${u.email}`
        .toLowerCase()
        .includes(term)
    );
  }

  selectUser(user: User): void {
    this.selectedUser = user;
    this.showUserModal = true;
    this.search = '';
    this.filteredAvail = [];
  }

  closeModal(): void {
    this.showUserModal = false;
    this.selectedUser = undefined;
  }

  addMember(user: User, role: string): void {
    let payload = {
      userID: user.id,
      teamID: this.teamID,
      role: role
    };
    this.teamSvc.newMember(payload).subscribe({
      next: (resp: MemberAPI) => {
        this.getMembers();
      },
      error: (err: Error) => {
        console.log(err.details);
      }
    });
    this.closeModal();
  }

  @HostListener('document:keydown.escape')
  handleEscape() {
    // Close modal first (priority)
    if (this.showUserModal) {
      this.closeModal();
      return;
    }

    // Then close sidebar
    if (this.sidebarOpen) {
      this.sidebarOpen = false;
      this.cd.detectChanges();
    }
  }

}
