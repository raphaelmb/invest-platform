-- +goose Up
-- +goose StatementBegin
CREATE TABLE wallet(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wallet;
-- +goose StatementEnd
