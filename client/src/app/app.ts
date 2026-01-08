import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { Auth } from './services/auth';
import { Error } from './models/error.model';
import { ChangeDetectorRef } from '@angular/core';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App implements OnInit {
  protected readonly title = signal('angular');

  log : boolean = false;

  private auth = inject(Auth);
  private cd = inject(ChangeDetectorRef);

  ngOnInit(): void {
    
    this.auth.isLogged.subscribe((status) => {
      if (status) {
        this.auth.testToken().subscribe({
          next: () => {
            this.log = true;
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
