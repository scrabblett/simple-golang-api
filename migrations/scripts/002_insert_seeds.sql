-- +goose Up
INSERT INTO age_groups (group_age, group_name)
VALUES (6, '6+');

INSERT INTO age_groups (group_age, group_name)
VALUES (12, '12+');

INSERT INTO age_groups (group_age, group_name)
VALUES (18, '18+');

INSERT INTO age_groups (group_age, group_name)
VALUES (16, '16+');

INSERT INTO library_user (last_name, first_name, patronymic, birth_date, created_at, edited_at, removed_at)
VALUES ('admin', 'admin', 'admin', '01-01-1900', current_date, current_date, null);

INSERT INTO users_info (login, password, salt, user_id)
SELECT 'admin',
       E'\\xD4D452001FDDD95B1921548C6536941B9C631158A33F981629A59590D6506D62',
       E'\\xB45E72C2EBC963CC',
       id
FROM library_user
WHERE first_name = 'admin';