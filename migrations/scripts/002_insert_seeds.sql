-- +goose Up
INSERT INTO public.age_groups (group_age, group_name)
VALUES (6, '6+');

INSERT INTO public.age_groups (group_age, group_name)
VALUES (12, '12+');

INSERT INTO public.age_groups (group_age, group_name)
VALUES (18, '18+');

INSERT INTO public.age_groups (group_age, group_name)
VALUES (16, '16+');

INSERT INTO library_user (last_name, first_name, patronymic, birth_date, created_at, edited_at, removed_at)
VALUES ('admin', 'admin', 'admin', '01-01-1900', current_date, current_date, null);

INSERT INTO users_info (login, password, salt, user_id)
SELECT 'admin', '2e735fd45433813b166a4e7379d52948fd65f2f0215fcba32b26c60abe1f06e9', '5cf5546e4e19b43apass', id
FROM public.library_user
WHERE first_name = 'admin';

INSERT INTO users_tokens (token, expired_at, user_id)
select '12345', '01-01-3000', user_id
from users_info
where login = 'admin';