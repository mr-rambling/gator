-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched TIMESTAMPTZ;

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetched;