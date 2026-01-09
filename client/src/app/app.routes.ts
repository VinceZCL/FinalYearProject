import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { PageNotFound } from './components/page-not-found/page-not-found';
import { authGuard } from './guards/auth-guard';
import { Home } from './components/home/home';
import { Profile } from './components/profile/profile';

export const routes: Routes = [
    {path: "login", component: Login},
    {path: "home", component: Home, canActivate: [authGuard]},
    {path: "profile", component: Profile, canActivate: [authGuard]},
    // TODO setup more sub-routes for profile (teams, etc.)

    // TODO setup admin page
    {path: "", redirectTo: "home", pathMatch: "full"},
    {path: "**", component: PageNotFound}
];