export interface NewComment {
    userID: number;
    checkinID: number;
    teamID: number;
    item: string;
}

export interface Comment {
    id: number;
    userID: number;
    username: string;
    item: string;
    checkinID: number;
    teamID: number;
    created_at: string;
}

export interface NewCommentAPI {
    status: string;
}