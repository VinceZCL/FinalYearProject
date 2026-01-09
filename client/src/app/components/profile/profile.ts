import { Component, inject, OnInit } from '@angular/core';
import { Auth } from '../../services/auth';
import { Claims } from '../../models/auth.model';

@Component({
  selector: 'app-profile',
  imports: [],
  templateUrl: './profile.html',
  styleUrl: './profile.css',
})
export class Profile implements OnInit {

  private auth = inject(Auth);
  claims! : Claims;

  ngOnInit() : void {
    this.auth.getClaims().subscribe({
      next: (claim : Claims) => {
        this.claims = claim;
      }
    });
  }



}
