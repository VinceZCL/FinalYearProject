export interface CheckIn {
    id: number;
    type: string;
    item: string;
    jira: string;
    visibility: string;
    teamID: number;
    userID: number;
    username: string;
    team_name: string;
    created_at: string;
}

export interface CheckInsAPI {
    status: string;
    checkIns: CheckIns | null;
}

export interface CheckInAPI {
    status: string;
    checkIn: CheckIn;
}

export interface CheckIns {
    userID: number;
    username: string;
    created_at: string;
    yesterday: CheckIn[];
    today: CheckIn[];
    blockers: CheckIn[];
}

export interface NewCheckIns {
    userID: number;
    type: 'yesterday' | 'today' | 'blockers';
    item?: string;
    jira?: string;
    visibility: string;
    teamID?: number;
}