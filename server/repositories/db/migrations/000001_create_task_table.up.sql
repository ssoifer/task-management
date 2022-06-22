create table if not exists task(
    "id" uuid primary key default gen_random_uuid(),
    "title" varchar,
    "content" varchar,
    "views" int,
    "timestamp" timestamp default current_timestamp
    );