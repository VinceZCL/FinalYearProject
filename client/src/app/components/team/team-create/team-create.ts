import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { AuthService } from '../../../services/auth';
import { Claims } from '../../../models/auth.model';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { TeamService } from '../../../services/team';
import { NgClass } from '@angular/common';
import { TeamCreatedAPI } from '../../../models/team.model';
import { Router } from '@angular/router';
import { Error } from '../../../models/error.model';

@Component({
  selector: 'app-team-create',
  imports: [ReactiveFormsModule, NgClass],
  templateUrl: './team-create.html',
  styleUrl: './team-create.css',
})
export class TeamCreate implements OnInit {

  private fb = inject(FormBuilder);
  private auth = inject(AuthService);
  private teamSvc = inject(TeamService);
  private router = inject(Router);
  private cd = inject(ChangeDetectorRef);

  uid!: number;
  form!: FormGroup;
  error: string = "";
  createFail: boolean = false;

  ngOnInit(): void {

    this.form = this.fb.group({
      name: ["", Validators.required]
    });

    this.auth.getClaims().subscribe({
      next: (claims: Claims) => {
        this.uid = claims.userID;
      }
    });
  }

  onSubmit(): void {
    if (this.form.invalid) {
      this.error = "Team Name is required";
      return;
    }
    this.error = "";

    let payload = {
      name: this.form.value.name,
      creatorID: this.uid
    };

    this.teamSvc.newTeam(payload).subscribe({
      next: (resp: TeamCreatedAPI) => {
        this.error = "";
        this.createFail = false;
        this.router.navigate(["/teams"]);
      },
      error: (err: Error) => {
        this.createFail = true;
        this.markName();
        this.error = err.details;
        console.log(err.details);
        this.cd.detectChanges();
      }
    })
  }

  invalidName(): boolean {
    let control = this.form.get("name");
    return !!(control && control.invalid && (control.dirty || control.touched) || this.createFail);
  }

  markName(): void {
    this.form.get("name")?.markAsTouched();
    this.form.get("name")?.markAsDirty();
  }

}
