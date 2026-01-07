export class Claims {

    userID: number;
    username: string;
    email: string;
    type: string;
    iss: string;
    exp: number;
    iat: number;

    constructor(claims: Claims) {
        this.userID = claims.userID;
        this.username = claims.username;
        this.email = claims.email;
        this.type = claims.type;
        this.iss = claims.iss;
        this.exp = claims.exp;
        this.iat = claims.iat;
    }
}

export interface IClaims {
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
    claims: IClaims;
}

export interface Login {
    status: string;
    token: string;
}