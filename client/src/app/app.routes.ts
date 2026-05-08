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
import { TeamDashboard } from './components/team/team-dashboard/team-dashboard';

export const routes: Routes = [
    {path: "login", component: Login, title: "Login"},

    // {path: "home", component: Home, canActivate: [authGuard]},
    {path: "dashboard", component: Dashboard, canActivate: [authGuard], title: "Personal Dashboard"},
    
    {path: "teams", canActivate: [authGuard], children: [
        {path: "", component: Team, title: "Teams"},
        {path: "create", component: TeamCreate, title: "Create new Team"},
        {path: ":teamID", component: TeamDashboard, title: "Team"},
    ]},

    {path: "profile", component: Profile, canActivate: [authGuard], title: "Profile"},

    {path: "admin", canActivate: [authGuard, adminGuard], children: [
        {path: "users", component: UserList, title: "Users"},
        {path: "users/create", component: UserCreate, title: "Create new User"},
    ]},

    // {path: "admin", component: Admin, canActivate: [adminGuard], children: [
        // {path: "admin/users", component: UserList, canActivate: [adminGuard]},
        // {path: "admin/users/create", component: UserCreate, canActivate: [adminGuard]},
    // ]},

    {path: "", redirectTo: "dashboard", pathMatch: "full"},
    {path: "**", component: PageNotFound, title: "Page Not Found"}
];