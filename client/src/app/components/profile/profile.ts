import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { User } from '../../models/user.model';
import { UserService } from '../../services/user';
import { ActivatedRoute } from '@angular/router';
import { Error } from '../../models/error.model';

@Component({
  selector: 'app-profile',
  imports: [],
  templateUrl: './profile.html',
  styleUrl: './profile.css',
})
export class Profile implements OnInit {

  private userSvc = inject(UserService);
  private cd = inject(ChangeDetectorRef);
  private route = inject(ActivatedRoute);
  uid!: number;
  user!: User;

  ngOnInit(): void {

    this.route.queryParams.subscribe(
      (params) => {
        let param = params["id"];
        this.uid = param !== null ? parseInt(param) : 0;
        this.update();
      }
    )

  }

  update(): void {
    this.userSvc.getUser(this.uid).subscribe({
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
