import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AuthService } from './services/auth';
import { Error } from './models/error.model';
import { ChangeDetectorRef } from '@angular/core';
import { AuthApi } from './models/auth.model';

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
  private cd = inject(ChangeDetectorRef);

  ngOnInit(): void {
    
    this.auth.getStatus().subscribe((status) => {
      if (status) {
        this.auth.testToken().subscribe({
          next: (resp : AuthApi) => {
            this.log = true;
            this.uid = resp.claims.userID;
            if (resp.claims.type == "admin") {
              this.admin = true;
            }
            this.cd.detectChanges();
          },
          error: (err : Error) => {
            console.log(err.details);
            this.logout();
            this.log = false;
            this.cd.detectChanges();
          }
        })
      }
    });
  }

  logout() : void {
    this.auth.logout();
  }

}
