import {  inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth';
import { catchError, map, of } from 'rxjs';

export const authGuard: CanActivateFn = (route, state) => {
  const auth = inject(AuthService);
  const router = inject(Router);

  if (auth.hasToken()) {
    return auth.testToken().pipe(
      map(() => {
        return true; // Token is valid, allow access
      }),
      catchError((err: Error) => {
        auth.logout(); // Log out on error
        router.navigate(["/login"]); // Redirect to login page
        return of(false); // Block the route
      })
    );
  } else {
    router.navigate(["/login"]);
    return of(false); // Block the route if no token
  }
};

