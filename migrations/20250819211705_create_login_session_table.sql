-- +goose Up
-- +goose StatementBegin
create table login_session
(
    token       text    not null
        constraint login_session_pk
            primary key,
    time_issued text,
    user_id     integer not null
        constraint login_session_user_id_fk
            references user
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE login_session
-- +goose StatementEnd
