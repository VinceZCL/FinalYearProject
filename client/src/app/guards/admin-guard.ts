import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth';
import { map, tap } from 'rxjs';

export const adminGuard: CanActivateFn = (route, state) => {
  const auth = inject(AuthService);
  const router = inject(Router);

  return auth.claim$.pipe(
    map(claims => claims?.type === "admin"),
    tap(isAdmin => {
      if (!isAdmin) {
        router.navigate(['/dashboard']);
      }
    })
  );


};
