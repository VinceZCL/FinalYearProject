import { Component, inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-dashboard',
  imports: [],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css',
})
export class Dashboard implements OnInit {

  readonly MAX_FIELDS = 2;

  form!: FormGroup;

  private fb = inject(FormBuilder);

  ngOnInit(): void {

    




  }

}
