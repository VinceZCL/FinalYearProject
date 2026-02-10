export interface Team {
    id: number;
    name: string;
    creatorID: number;
    creator_name: string;
}

export interface Member {
    userID: number;
    name: string;
    email: string;
    teamID: number;
    team_name: string;
    role: "admin" | "member";
}

export interface TeamAPI {
    status: string;
    team: Team;
}

export interface MembersAPI {
    status: string;
    members: Member[];
}

export interface TeamCreatedAPI {
    success: string;
    team: Team;
    member: Member;
}