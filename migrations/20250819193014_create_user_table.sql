-- +goose Up
-- +goose StatementBegin
create table user
(
    id         integer not null
        constraint user_pk
            primary key autoincrement,
    first_name text    not null,
    last_name  integer not null,
    email      text    not null,
    picture    text    not null,
    sub        text    not null,
    issuer_id  integer not null
        constraint user__oauth_issuer_id_fk
            references oauth_issuer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user
-- +goose StatementEnd
