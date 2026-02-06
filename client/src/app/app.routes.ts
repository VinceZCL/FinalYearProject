import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { PageNotFound } from './components/page-not-found/page-not-found';
import { authGuard } from './guards/auth-guard';
import { Profile } from './components/profile/profile';
import { UserList } from './components/admin/user-list/user-list';
import { UserCreate } from './components/admin/user-create/user-create';
import { adminGuard } from './guards/admin-guard';
import { Dashboard } from './components/dashboard/dashboard';
import { Team } from './components/team/team';
import { TeamCreate } from './components/team/team-create/team-create';

export const routes: Routes = [
    {path: "login", component: Login},

    // {path: "home", component: Home, canActivate: [authGuard]},

    {path: "dashboard", component: Dashboard, canActivate: [authGuard]},    
    
    {path: "teams", children: [
        {path: "", component: Team, canActivate: [authGuard]},
        {path: "create", component: TeamCreate, canActivate: [authGuard]},
    ]},

    {path: "profile", component: Profile, canActivate: [authGuard]},

    {path: "admin", children: [
        {path: "users", component: UserList, canActivate: [adminGuard]},
        {path: "users/create", component: UserCreate, canActivate: [adminGuard]},
    ]},

    // {path: "admin", component: Admin, canActivate: [adminGuard], children: [
        // {path: "admin/users", component: UserList, canActivate: [adminGuard]},
        // {path: "admin/users/create", component: UserCreate, canActivate: [adminGuard]},
    // ]},

    {path: "", redirectTo: "dashboard", pathMatch: "full"},
    {path: "**", component: PageNotFound}
];