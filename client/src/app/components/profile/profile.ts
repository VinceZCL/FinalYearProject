import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { User } from '../../models/user.model';
import { UserService } from '../../services/user';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { Error } from '../../models/error.model';
import { effect } from '@angular/core';
import { AuthService } from '../../services/auth';
import { TeamService } from '../../services/team';
import { Member } from '../../models/team.model';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgClass, UpperCasePipe } from '@angular/common';

@Component({
  selector: 'app-profile',
  imports: [RouterLink, ReactiveFormsModule, NgClass, UpperCasePipe],
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
  updateFail: boolean = false;
  error: string = "";

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
      email: [{ value: this.user.email, disabled: true }, [Validators.required, Validators.email]],
      current_password: ["", Validators.required],
      new_password: [""]
    });
  }

  onSubmit(): void {

    this.updateFail = false;
    if (this.form.invalid) {
      this.updateFail = true;
      this.markAllFields();

      let nameErr = this.getForm("name").errors;
      let emailErr = this.getForm("email").errors;
      let passwordErr = this.getForm("current_password").errors;

      if (nameErr) {
        this.error = "Name is required";
        return;
      }

      if (emailErr) {
        if (emailErr["required"]) {
          this.error = "Email is required";
        }
        if (emailErr["email"]) {
          this.error = "Email format is invalid";
        }
        return;
      }
      if (passwordErr) {
        this.error = "Password is required";
        return;
      }
    }
    this.error = "";

    this.userSvc.updateUser(this.uid, {userID: this.uid, ...this.form.getRawValue()}).subscribe({
      next: (val: User) => {
        this.update();
        this.edit = false;
      },
      error: (err: Error) => {
        this.updateFail = true;
        this.markAllFields();
        this.error = err.details;
        console.log(err.details);
        this.cd.detectChanges();
      }
    });
  }

  private getForm(controlName: string): FormControl {
    return this.form.get(controlName) as FormControl;
  }

  invalidInp(controlName: string): boolean {
    let control = this.getForm(controlName);
    return !!(control && control.invalid && (control.dirty || control.touched) || this.updateFail);
  }

  private markAllFields(): void {
    Object.values(this.form.controls).forEach(control => {
      control.markAsTouched();
      control.markAsDirty();
    })
  }

}
