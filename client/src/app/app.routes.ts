import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { PageNotFound } from './components/page-not-found/page-not-found';
import { authGuard } from './guards/auth-guard';
import { Home } from './components/home/home';

export const routes: Routes = [
    {path: "login", component: Login},
    {path: "home", component: Home, canActivate: [authGuard]},
    {path: "", redirectTo: "home", pathMatch: "full"},
    {path: "**", component: PageNotFound}
];

// TODO setup routes for login