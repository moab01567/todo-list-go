-- +goose Up
-- +goose StatementBegin
create table oauth_issuer
(
    id integer not null constraint oauth_issuer_pk primary key,
    issuer TEXT not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE oauth_issuer
-- +goose StatementEnd
