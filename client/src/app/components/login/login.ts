import { NgClass } from '@angular/common';
import { ChangeDetectorRef, Component, inject } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth';
import { Error } from '../../models/error.model';

@Component({
  selector: 'app-login',
  imports: [ReactiveFormsModule, NgClass],
  templateUrl: './login.html',
  styleUrl: './login.css',
})
export class Login {

  private fb = inject(FormBuilder);
  private auth = inject(AuthService);
  private router = inject(Router);
  private cd = inject(ChangeDetectorRef);

  form : FormGroup;
  error! : string;
  loginFail! : boolean;

  constructor() {
    this.form = this.fb.group({
      email: ["", [Validators.required, Validators.email]],
      password: ["", Validators.required]
    });
  }

  onSubmit() {
    this.loginFail = false;
    if (this.form.invalid) {
      this.markAllFields();
      
      let emailErr = this.getForm("email").errors;
      let passwordErr = this.getForm("password").errors;

      if (emailErr) {
        if (emailErr['required']) {
          this.error = 'Email is required';
        }
        if (emailErr['email']) {
          this.error = 'Email format is invalid';
        }
        return;
      }
      if (passwordErr) {
        if (passwordErr['required']) {
          this.error = 'Password is required';
        }
        return;
      }
    };
    this.error = "";

    this.auth.login(this.form.value).subscribe({
      next: () => {
        this.error = "";
        this.loginFail = false;
        this.router.navigate(["/dashboard"]);
      },
      error: (err : Error) => {
        this.loginFail = true;
        this.markAllFields();
        this.error = err.details;
        console.log(err.details);
        this.cd.detectChanges();
      }
    });

  }

  private getForm(controlName: string) : FormControl {
    return this.form.get(controlName) as FormControl;
  }

  invalidInp(controlName: string) : boolean {
    let control = this.getForm(controlName);
    return !!(control && control.invalid && (control.dirty || control.touched) || this.loginFail);
  }

  private markAllFields() : void {
    Object.values(this.form.controls).forEach(control => {
      control.markAsTouched();
      control.markAsDirty();
    })
  }

}
