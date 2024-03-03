-- +goose Up
-- +goose StatementBegin
CREATE TABLE data
    (login varchar(128) NOT NULL,
     type varchar(128) NOT NULL,
     name varchar(128) NOT NULL,
     meta bytea,
     data bytea,
     UNIQUE(login, type, name));
CREATE INDEX login_index ON data(login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE data;
-- +goose StatementEnd
