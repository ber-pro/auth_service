-- +goose Up
-- +goose StatementBegin
    CREATE TABLE users (
        id serial PRIMARY KEY,
        username varchar(50) NOT NULL,
        email varchar(100) NOT NULL,
        password varchar(255) NOT NULL,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE users;
-- +goose StatementEnd
