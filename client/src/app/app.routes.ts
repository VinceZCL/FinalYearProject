import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { PageNotFound } from './components/page-not-found/page-not-found';
import { authGuard } from './guards/auth-guard';
import { Home } from './components/home/home';
import { Profile } from './components/profile/profile';
import { Admin } from './components/admin/admin';
import { UserList } from './components/admin/user-list/user-list';
import { UserCreate } from './components/admin/user-create/user-create';

export const routes: Routes = [
    {path: "login", component: Login},
    {path: "home", component: Home, canActivate: [authGuard]},
    {path: "profile", component: Profile, canActivate: [authGuard]},
    // TODO setup more sub-routes for profile (teams, etc.)

    // TODO setup admin page

    {path: "admin", component: Admin, canActivate: [authGuard], children: [
        {path: "users", component: UserList, canActivate: [authGuard]},
        {path: "users/create", component: UserCreate, canActivate: [authGuard]}
    ]},

    {path: "", redirectTo: "home", pathMatch: "full"},
    {path: "**", component: PageNotFound}
];