-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
    (login varchar(128) PRIMARY KEY NOT NULL,
     password text NOT NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
