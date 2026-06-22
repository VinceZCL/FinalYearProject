--liquibase formatted sql

--changeset fyp-scrum:001-fyp-scrum-users-table-deleted-index runOnChange:false
--comment: Create fyp scrum users table deleted index
CREATE INDEX idx_fyp_scrum_users_deleted_at ON fyp_scrum_users(deleted_at);
--rollback DROP INDEX idx_fyp_scrum_users_deleted_at;

--changeset fyp-scrum:002-fyp-scrum-teams-table-deleted-index runOnChange:false
--comment: Create fyp scrum teams table deleted index
CREATE INDEX idx_fyp_scrum_teams_deleted_at ON fyp_scrum_teams(deleted_at);
--rollback DROP INDEX idx_fyp_scrum_teams_deleted_at;

--changeset fyp-scrum:003-fyp-scrum-teams-table-fk-constraint runOnChange:false
--comment: Create fyp scrum teams table fk constraint
ALTER TABLE fyp_scrum_teams 
    ADD CONSTRAINT fk_fyp_scrum_teams_creator
        FOREIGN KEY (creator_id)
        REFERENCES fyp_scrum_users(id)
        ON UPDATE CASCADE
        ON DELETE SET NULL;
--rollback ALTER TABLE fyp_scrum_teams DROP CONSTRAINT fk_fyp_scrum_team_name_creator;

--changeset fyp-scrum:004-fyp-scrum-teams-table-unique-constraint runOnChange:false
--comment: Create fyp scrum teams table unique constraint
ALTER TABLE fyp_scrum_teams 
    ADD CONSTRAINT uq_fyp_scrum_team_name_creator 
        UNIQUE (name, creator_id);
--rollback ALTER TABLE fyp_scrum_teams DROP CONSTRAINT uq_fyp_scrum_team_name_creator;

--changeset fyp-scrum:005-fyp-scrum-user-teams-table-fk-user-constraint runOnChange:false
--comment: Create fyp scrum teams table fk user constraint
ALTER TABLE fyp_scrum_user_teams
    ADD CONSTRAINT fk_fyp_scrum_user_teams_user
        FOREIGN KEY (user_id)
        REFERENCES fyp_scrum_users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;
--rollback ALTER TABLE fyp_scrum_user_teams DROP CONSTRAINT fk_fyp_scrum_user_teams_user;

--changeset fyp-scrum:006-fyp-scrum-user-teams-table-fk-team-constraint runOnChange:false
--comment: Create fyp scrum teams table fk team constraint
ALTER TABLE fyp_scrum_user_teams
    ADD CONSTRAINT fk_fyp_scrum_user_teams_team
        FOREIGN KEY (team_id)
        REFERENCES fyp_scrum_teams(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;
--rollback ALTER TABLE fyp_scrum_user_teams DROP CONSTRAINT fk_fyp_scrum_user_teams_team;

--changeset fyp-scrum:007-fyp-scrum-checkins-table-fk-user-constraint runOnChange:false
--comment: Create fyp scrum checkin table fk user constraint
ALTER TABLE fyp_scrum_checkins
    ADD CONSTRAINT fk_fyp_scrum_checkins_user
        FOREIGN KEY (user_id)
        REFERENCES fyp_scrum_users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;
--rollback ALTER TABLE fyp_scrum_checkins DROP CONSTRAINT fk_fyp_scrum_checkins_user;

--changeset fyp-scrum:008-fyp-scrum-checkins-table-fk-team-constraint runOnChange:false
--comment: Create fyp scrum checkin table fk team constraint
ALTER TABLE fyp_scrum_checkins
    ADD CONSTRAINT fk_fyp_scrum_checkins_team
        FOREIGN KEY (team_id)
        REFERENCES fyp_scrum_teams(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;
--rollback ALTER TABLE fyp_scrum_checkins DROP CONSTRAINT fk_fyp_scrum_checkins_team;

--changeset fyp-scrum:009-fyp-scrum-checkins-table-user-index runOnChange:false
--comment: Create fyp scrum checkin table user id index
CREATE INDEX idx_fyp_scrum_checkins_user_id ON fyp_scrum_checkins(user_id);
--rollback DROP INDEX idx_fyp_scrum_checkins_user_id;

--changeset fyp-scrum:010-fyp-scrum-checkins-table-team-index runOnChange:false
--comment: Create fyp scrum checkin table team id index
CREATE INDEX idx_fyp_scrum_checkins_team_id ON fyp_scrum_checkins(team_id);
--rollback DROP INDEX idx_fyp_scrum_checkins_team_id;

--changeset fyp-scrum:011-fyp-scrum-checkins-table-deleted-index runOnChange:false
--comment: Create fyp scrum checkin table deleted index
CREATE INDEX idx_fyp_scrum_checkins_deleted_at ON fyp_scrum_checkins(deleted_at);
--rollback DROP INDEX idx_fyp_scrum_checkins_deleted_at;

--changeset fyp-scrum:012-fyp-scrum-comments-table-fk-user-constraint runOnChange:false
--comment: Create fyp scrum comments table fk user constraint
ALTER TABLE fyp_scrum_comments
    ADD CONSTRAINT fk_fyp_scrum_comments_user
        FOREIGN KEY (user_id)
        REFERENCES fyp_scrum_users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;
--rollback ALTER TABLE fyp_scrum_comments DROP CONSTRAINT fk_fyp_scrum_comments_user;

--changeset fyp-scrum:013-fyp-scrum-comments-table-fk-checkin-user-constraint runOnChange:false
--comment: Create fyp scrum comments table fk checkin user constraint
ALTER TABLE fyp_scrum_comments
    ADD CONSTRAINT fk_fyp_scrum_comments_checkin
        FOREIGN KEY (checkin_id)
        REFERENCES fyp_scrum_checkins(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;
--rollback ALTER TABLE fyp_scrum_comments DROP CONSTRAINT fk_fyp_scrum_comments_checkin_user;

--changeset fyp-scrum:014-fyp-scrum-comments-table-fk-team-constraint runOnChange:false
--comment: Create fyp scrum comments table fk team constraint
ALTER TABLE fyp_scrum_comments
    ADD CONSTRAINT fk_fyp_scrum_comments_team
        FOREIGN KEY (team_id)
        REFERENCES fyp_scrum_teams(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;
--rollback ALTER TABLE fyp_scrum_comments DROP CONSTRAINT fk_fyp_scrum_comments_team;

--changeset fyp-scrum:015-fyp-scrum-comments-table-user-index runOnChange:false
--comment: Create fyp scrum comments table user id index
CREATE INDEX idx_fyp_scrum_comments_user_id
    ON fyp_scrum_comments(user_id);
--rollback DROP INDEX idx_fyp_scrum_comments_user_id;

--changeset fyp-scrum:016-fyp-scrum-comments-table-checkin-user-index runOnChange:false
--comment: Create fyp scrum comments table checkin user id index
CREATE INDEX idx_fyp_scrum_comments_checkin_id
    ON fyp_scrum_comments(checkin_id);
--rollback DROP INDEX idx_fyp_scrum_comments_checkin_user_id;

--changeset fyp-scrum:017-fyp-scrum-comments-table-team-index runOnChange:false
--comment: Create fyp scrum comments table team id index
CREATE INDEX idx_fyp_scrum_comments_team_id
    ON fyp_scrum_comments(team_id);
--rollback DROP INDEX idx_fyp_scrum_comments_team_id;

--changeset fyp-scrum:018-fyp-scrum-comments-table-deleted-index runOnChange:false
--comment: Create fyp scrum comments table deleted index
CREATE INDEX idx_fyp_scrum_comments_deleted_at
    ON fyp_scrum_comments(deleted_at);
--rollback DROP INDEX idx_fyp_scrum_comments_deleted_at;