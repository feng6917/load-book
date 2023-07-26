-- migrate -verbose -database "postgres://postgres:123456@127.0.0.1:5432/qqyyds?sslmode=disable" -path "migration" up
create table "categories"
(
    "id"        serial primary key,
    "name"      varchar(256),
    "desc"      varchar(512),
    "create_at" TIMESTAMPTZ NOT NULL,
    "update_at" TIMESTAMPTZ
);

