import { ChangeDetectorRef, Component, inject } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../../../services/user';
import { UserAPI } from '../../../models/user.model';
import { Error } from '../../../models/error.model';
import { NgClass } from '@angular/common';

@Component({
  selector: 'app-user-create',
  imports: [ReactiveFormsModule, NgClass],
  templateUrl: './user-create.html',
  styleUrl: './user-create.css',
})
export class UserCreate {

  private fb = inject(FormBuilder);
  private userSvc = inject(UserService);
  private router = inject(Router);
  private cd = inject(ChangeDetectorRef);

  form: FormGroup;
  error: string = "";
  createFail: boolean = false;

  constructor() {
    this.form = this.fb.group({
      name: ["", Validators.required],
      email: ["", [Validators.required, Validators.email]],
      password: ["", Validators.required],
      type: ["user", Validators.required]
    });
  }

  onSubmit() {
    this.createFail = false;
    if (this.form.invalid) {
      this.markAllFields();

      let nameErr = this.getForm("name").errors;
      let emailErr = this.getForm("email").errors;
      let passwordErr = this.getForm("password").errors;
      let typeErr = this.getForm("type").errors;

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
      if (typeErr) {
        this.error = "User Type is required";
        return;
      }
    };
    this.error = "";

    this.userSvc.newUser(this.form.value).subscribe({

      next: (resp: UserAPI) => {
        this.error = "";
        this.createFail = false;
        let uid = resp.user.id;
        this.router.navigate(["/profile"], {queryParams: {id: uid}});
      },
      error: (err: Error) => {
        this.createFail = true;
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
    return !!(control && control.invalid && (control.dirty || control.touched) || this.createFail);
  }

  private markAllFields(): void {
    Object.values(this.form.controls).forEach(control => {
      control.markAsTouched();
      control.markAsDirty();
    })
  }

}
