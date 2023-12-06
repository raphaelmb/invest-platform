-- +goose Up
-- +goose StatementBegin
CREATE TABLE assets (
    id UUID PRIMARY KEY,
    symbol TEXT NOT NULL,
    price NUMERIC NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE assets;
-- +goose StatementEnd
