-- +goose Up
-- +goose StatementBegin
CREATE TYPE ORDER_TYPE AS ENUM ('BUY', 'SELL');
CREATE TYPE ORDER_STATUS AS ENUM ('PENDING', 'OPEN', 'CLOSED', 'FAILED');
CREATE TABLE "order"(
    id UUID PRIMARY KEY,
    type ORDER_TYPE NOT NULL,
    status ORDER_STATUS NOT NULL,
    shares INTEGER NOT NULL,
    price NUMERIC NOT NULL,
    partial INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    wallet_id UUID NOT NULL REFERENCES wallet(id) ON DELETE CASCADE,
    asset_id UUID NOT NULL REFERENCES assets(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "order";
DROP TYPE ORDER_TYPE;
DROP TYPE ORDER_STATUS;
-- +goose StatementEnd
