-- +goose Up
-- +goose StatementBegin
create table todo
(
    id      integer               not null
        constraint todo_pk
            primary key autoincrement,
    name    text                  not null,
    done    BOOLEAN default 0,
    user_id integer               not null
        constraint todo_user_id_fk
            references user
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todo
-- +goose StatementEnd
