import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { PageNotFound } from './components/page-not-found/page-not-found';
import { authGuard } from './guards/auth-guard';
import { Home } from './components/home/home';
import { Profile } from './components/profile/profile';
import { Admin } from './components/admin/admin';
import { UserList } from './components/admin/user-list/user-list';
import { UserCreate } from './components/admin/user-create/user-create';
import { adminGuard } from './guards/admin-guard';
import { Dashboard } from './components/dashboard/dashboard';
import { Team } from './components/team/team';

export const routes: Routes = [
    {path: "login", component: Login},
    // {path: "home", component: Home, canActivate: [authGuard]},
    {path: "dashboard", component: Dashboard, canActivate: [authGuard]},
    {path: "teams", component: Team, canActivate: [authGuard]},
    
    // TODO teams => { list, dashboard => { members } }

    {path: "profile", component: Profile, canActivate: [authGuard]},
    // TODO setup more sub-routes for profile (teams, etc.)

    // {path: "admin", component: Admin, canActivate: [adminGuard], children: [
        {path: "admin/users", component: UserList, canActivate: [adminGuard]},
        {path: "admin/users/create", component: UserCreate, canActivate: [adminGuard]},
    // ]},

    {path: "", redirectTo: "dashboard", pathMatch: "full"},
    {path: "**", component: PageNotFound}
];