import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class Auth {

  private url: string = "http://localhost:8080/api/auth";
  private http = inject(HttpClient)

}
