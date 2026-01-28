import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { User } from '../../models/user.model';
import { UserService } from '../../services/user';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { Error } from '../../models/error.model';
import { effect } from '@angular/core';
import { AuthService } from '../../services/auth';
import { AuthApi } from '../../models/auth.model';

@Component({
  selector: 'app-profile',
  imports: [RouterLink],
  templateUrl: './profile.html',
  styleUrl: './profile.css',
})
export class Profile implements OnInit {

  private userSvc = inject(UserService);
  private cd = inject(ChangeDetectorRef);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private auth = inject(AuthService);
  uid!: number;
  user!: User;
  admin!: boolean;
  showback: boolean = false;

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
        this.uid = param !== null ? parseInt(param) : 0;
        this.update();
        this.checkDelete();
      }
    );

    this.auth.testToken().subscribe({
      next: (resp : AuthApi) => {
        this.admin = (resp.claims.type == "admin" && resp.claims.userID != this.uid) ? true : false;
        // if (resp.claims.userID == this.uid) {
        //   this.admin = false;
        // }
        this.cd.detectChanges();
      }
    })

  }

  update(): void {
    this.userSvc.getUser(this.uid).subscribe({
      next: (resp: User) => {
        this.user = resp;
        this.cd.detectChanges();
      },
      error: (err: Error) => {
        console.log(err.details);
        this.router.navigate(["/404"]);
      }
    })
  }

  checkDelete(): void {
    
    this.auth.testToken().subscribe({
      next: (resp : AuthApi) => {
        this.admin = (resp.claims.type == "admin" && resp.claims.userID != this.uid);
        this.cd.detectChanges();
      }
    })

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

}
