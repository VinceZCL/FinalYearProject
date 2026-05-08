--liquibase formatted sql

--changeset fyp-scrum:003-create-fyp-scrum-users-table runOnChange:false
--comment: Create fyp scrum users table
CREATE TABLE IF NOT EXISTS fyp_scrum_users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    type VARCHAR(50) DEFAULT 'user',
    status VARCHAR(50) DEFAULT 'active'
);
--rollback DROP TABLE IF EXISTS fyp_scrum_users;

--changeset fyp-scrum:004-create-fyp-scrum-teams-table runOnChange:false
--comment: Create fyp scrum teams table
CREATE TABLE IF NOT EXISTS fyp_scrum_teams (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    name VARCHAR(100) NOT NULL,
    creator_id BIGINT
);
--rollback DROP TABLE IF EXISTS fyp_scrum_teams;

--changeset fyp-scrum:005-create-fyp-user-teams-table runOnChange:false
--comment: Create fyp scrum user teams table
CREATE TABLE IF NOT EXISTS fyp_scrum_user_teams (
    user_id BIGINT NOT NULL,
    team_id BIGINT NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'member',

    PRIMARY KEY (user_id, team_id)
);
--rollback  DROP TABLE IF EXISTS fyp_scrum_user_teams;

--changeset fyp-scrum:006-create-fyp-scrum-checkins-table runOnChange:false
--comment: Create fyp scrum checkins table
CREATE TABLE IF NOT EXISTS fyp_scrum_checkins (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    user_id BIGINT NOT NULL,
    type VARCHAR(50) NOT NULL,
    item VARCHAR(255) NOT NULL,
    jira VARCHAR(100),
    visibility VARCHAR(10) NOT NULL DEFAULT 'all',
    team_id BIGINT
);
--rollback DROP TABLE IF EXISTS fyp_scrum_checkins;