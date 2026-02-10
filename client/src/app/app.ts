import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AuthService } from './services/auth';
import { ChangeDetectorRef } from '@angular/core';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App implements OnInit {
  protected readonly title = signal('angular');

  log! : boolean;
  admin! : boolean;
  uid! : number;

  private auth = inject(AuthService);

  ngOnInit(): void {
    this.auth.logged$.subscribe(isLogged => {
      this.log = isLogged;
    });

    this.auth.userID$.subscribe(uid => {
      this.uid = Number(uid);
    });

    this.auth.admin$.subscribe(isAdmin => {
      this.admin = isAdmin;
    })
  }

  logout() : void {
    this.auth.logout();
  }

}
