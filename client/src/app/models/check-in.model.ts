export interface CheckIn {
    id: number;
    type: string;
    item: string;
    jira: string;
    visibility: string;
    teamID: number;
    userID: number;
    username: string;
    created_at: string;
}

export interface CheckInsAPI {
    status: string;
    checkIns: checkins | null;
}

export interface CheckInAPI {
    status: string;
    checkIn: CheckIn;
}

interface checkins {
    userID: number;
    username: string;
    created_at: string;
    yesterday: CheckIn[];
    today: CheckIn[];
    blockers: CheckIn[];
}