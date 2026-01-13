import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { User } from '../../models/user.model';
import { UserService } from '../../services/user';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { Error } from '../../models/error.model';

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
  uid!: number;
  user!: User;
  showback: boolean = false;

  ngOnInit(): void {

    const nav = this.router.getCurrentNavigation();
    const state = nav?.extras.state ?? history.state;

    this.showback = state?.from === 'list';

    this.route.queryParams.subscribe(
      (params) => {
        let param = params["id"];
        this.uid = param !== null ? parseInt(param) : 0;
        this.update();
      }
    );

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

}
