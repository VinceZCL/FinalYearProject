import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { NavigationEnd, Router, RouterOutlet } from '@angular/router';
import { filter } from 'rxjs';

@Component({
  selector: 'app-admin',
  imports: [RouterOutlet],
  templateUrl: './admin.html',
  styleUrl: './admin.css',
})
export class Admin implements OnInit {

  private router = inject(Router);
  private cd = inject(ChangeDetectorRef)
  sub!: boolean;

  ngOnInit(): void {
    this.checkSub();

    // Listen to route changes (NavigationEnd event)
    this.router.events.pipe(
      filter((event) => event instanceof NavigationEnd)
    ).subscribe(() => {
      this.checkSub();
    });
  }
  
  checkSub() : void {
    let url = this.router.url;
    this.sub = url.includes('/users');
    this.cd.detectChanges();
  }

  navClick(route : string) : void {
    this.router.navigate([route]);
    this.checkSub();
  }

}
