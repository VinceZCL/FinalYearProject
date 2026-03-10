import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { User } from '../../models/user.model';
import { UserService } from '../../services/user';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { Error } from '../../models/error.model';
import { effect } from '@angular/core';
import { AuthService } from '../../services/auth';
import { TeamService } from '../../services/team';
import { Member } from '../../models/team.model';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
  selector: 'app-profile',
  imports: [RouterLink, ReactiveFormsModule],
  templateUrl: './profile.html',
  styleUrl: './profile.css',
})
export class Profile implements OnInit {

  private userSvc = inject(UserService);
  private cd = inject(ChangeDetectorRef);
  private teamSvc = inject(TeamService);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private auth = inject(AuthService);
  private fb = inject(FormBuilder);
  uid!: number;
  user!: User;
  admin!: boolean;
  self!: boolean;
  teams!: Member[];
  showback: boolean = false;
  edit: boolean = false;
  form!: FormGroup;

  // ✅ effect runs in injection context
  private navigationEffect = effect(() => {
    const nav = this.router.currentNavigation();
    const state = nav?.extras.state ?? history.state;

    this.showback = state?.from === 'list';
  });

  ngOnInit(): void {

    this.route.queryParams.subscribe(
      (params) => {
        let param = params["id"];
        this.uid = Number(param);
        this.checkDelete();
        this.update();
      }
    );
  }

  update(): void {
    this.userSvc.getUser(this.uid).subscribe({
      next: (resp: User) => {
        this.user = resp;
        this.getTeams();
        if (this.self) {
          this.initForm();
        }
        this.cd.detectChanges();
      },
      error: (err: Error) => {
        console.log(err.details);
        this.router.navigate(["/404"]);
      }
    })
  }

  checkDelete(): void {

    this.auth.claim$.subscribe(claims => {
      this.self = claims?.userID == this.uid;
      this.admin = claims?.type == "admin";
      this.cd.detectChanges();
    });

  }

  deactivate(id: number): void {
    this.userSvc.deactivateUser(id).subscribe({
      next: (resp: User) => {
        this.user = resp;
        this.cd.detectChanges();
      },
      error: (err: Error) => {
        console.log(err.details);
      }
    })
  }

  getTeams(): void {
    this.teamSvc.getOwnTeams(this.uid).subscribe({
      next: (resp: Member[]) => {
        this.teams = resp;
        this.teams.sort((a, b) => Number(b.role === "admin") - Number(a.role === "admin"));
        this.cd.detectChanges();
      },
      error: (err: Error) => {
        console.log(err.details);
      }
    })
  }

  initForm(): void {
    this.form = this.fb.group({
      name: [this.user.name, Validators.required],
      email: [this.user.email, [Validators.required, Validators.email]],
      current_password: ["", Validators.required],
      new_password: [""]
    });
  }

  onSubmit(): void {

    // TODO error checking

    this.userSvc.updateUser(this.uid, {userID: this.uid, ...this.form.value}).subscribe({
      next: (val: User) => {
        this.update();
        this.edit = false;
      },
      error: (err: Error) => {
        console.log(err.details);
      }
    });
  }

}
