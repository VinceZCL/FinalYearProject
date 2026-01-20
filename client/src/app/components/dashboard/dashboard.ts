import { Component, inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { AuthService } from '../../services/auth';
import { Claims } from '../../models/auth.model';

@Component({
  selector: 'app-dashboard',
  imports: [],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css',
})
export class Dashboard implements OnInit {

  readonly MAX_FIELDS = 2;

  form!: FormGroup;
  uid!: number;

  private fb = inject(FormBuilder);
  private auth = inject(AuthService);

  ngOnInit(): void {

    this.auth.getClaims().subscribe({
      next: (claims : Claims) => {
        this.uid = claims.userID;
        alert(`your uid: ${this.uid}`);
      }
    })

    // TODO trigger checkin.getDaily
    
  }

}
