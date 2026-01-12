export interface Claims {
    userID: number;
    username: string;
    email: string;
    type: string;
    iss: string;
    exp: number;
    iat: number;
}

export interface AuthApi {
    status: string;
    claims: Claims;
}

export interface Login {
    status: string;
    token: string;
}