import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { User } from '../../../models/user.model';
import { UserService } from '../../../services/user';
import { Error } from '../../../models/error.model';
import { UpperCasePipe } from '@angular/common';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-user-list',
  imports: [UpperCasePipe, RouterLink],
  templateUrl: './user-list.html',
  styleUrl: './user-list.css',
})
export class UserList implements OnInit {

  users! : User[];

  private userSvc = inject(UserService);
  private cd = inject(ChangeDetectorRef)

  ngOnInit(): void {
    this.userSvc.getUsers().subscribe({
      next: (resp : User[]) => {
        this.users = resp;
        this.cd.detectChanges();
      },
      error: (err : Error) => {
        console.log(err.details);
      }
    }
    )
  }

}
