export interface User {
    id: number;
    name: string;
    email: string;
    type: string;
}

export interface UsersAPI {
    status: string;
    users: User[];
}

export interface UserAPI {
    status: string;
    user: User;
}
