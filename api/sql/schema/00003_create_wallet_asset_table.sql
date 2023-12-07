-- +goose Up
-- +goose StatementBegin
CREATE TABLE wallet_asset(
    id UUID PRIMARY KEY,
    amount NUMERIC NOT NULL,
    shares INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    wallet_id UUID NOT NULL REFERENCES wallet(id) ON DELETE CASCADE,
    asset_id UUID NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    UNIQUE(wallet_id, asset_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wallet_asset;
-- +goose StatementEnd
