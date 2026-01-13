import { Component, inject } from '@angular/core';
import { UserService } from '../../../services/user';
import { AsyncPipe, UpperCasePipe } from '@angular/common';
import { RouterLink } from '@angular/router';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { combineLatest, map, startWith } from 'rxjs';

@Component({
  selector: 'app-user-list',
  imports: [UpperCasePipe, RouterLink, AsyncPipe, ReactiveFormsModule],
  templateUrl: './user-list.html',
  styleUrl: './user-list.css',
})
export class UserList {

  // users! : User[];

  // private userSvc = inject(UserService);
  // private cd = inject(ChangeDetectorRef)

  // ngOnInit(): void {
  //   this.userSvc.getUsers().subscribe({
  //     next: (resp : User[]) => {
  //       this.users = resp;
  //       this.cd.detectChanges();
  //     },
  //     error: (err : Error) => {
  //       console.log(err.details);
  //     }
  //   }
  //   )
  // }

  private userSvc = inject(UserService);

  // search input
  searchCtrl = new FormControl('');

  // source users stream
  users$ = this.userSvc.getUsers();

  // derived filtered users
  filteredUsers$ = combineLatest([
    this.users$,
    this.searchCtrl.valueChanges.pipe(startWith(''))
  ]).pipe(
    map(([users, search]) => {
      const term = search?.toLowerCase().trim() ?? '';
      return users.filter(u =>
        u.name.toLowerCase().includes(term) ||
        u.email.toLowerCase().includes(term) ||
        u.type.toLowerCase().includes(term)
      );
    })
  );

}
