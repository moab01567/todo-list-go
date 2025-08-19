-- +goose Up
-- +goose StatementBegin
INSERT into oauth_issuer VALUES (1,'google');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM oauth_issuer WHERE id = 1
-- +goose StatementEnd
