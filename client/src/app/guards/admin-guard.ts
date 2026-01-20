import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth';
import { catchError, map, of } from 'rxjs';
import { AuthApi } from '../models/auth.model';

export const adminGuard: CanActivateFn = (route, state) => {
  const auth = inject(AuthService);
  const router = inject(Router);

  if (auth.hasToken()) {
    return auth.testToken().pipe(
      map((resp : AuthApi) => {
        if (resp.claims.type == "admin") {
          return true; // Token is valid, allow access
        } else {
          router.navigate(["/dashboard"]);
          return false;
        }
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
