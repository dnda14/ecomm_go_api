-- +goose Up
-- +goose StatementBegin
create table if not exists products(
    id bigserial primary key,
    name text not null,
    price_in_center integer not null check (price_in_center >= 0),
    quantity integer not null default 0,
    create_at timestamptz not null default now()
);
-- +goose StatementEnd
    
-- +goose Down
-- +goose StatementBegin
drop table if exists products;
-- +goose StatementEnd
