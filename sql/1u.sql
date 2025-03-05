create extension if not exists citext;

CREATE TABLE
    IF NOT EXISTS public.users (
        id BIGSERIAL PRIMARY KEY,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        name TEXT NOT NULL,
        photo TEXT,
        role TEXT DEFAULT 'user' CHECK (role IN ('user', 'guide', 'admin')),
        active BOOLEAN DEFAULT TRUE,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );

-- create table
--     if not exists public.posts (
--         id bigserial primary key,
--         user_id bigint references public.users (id),
--         content text,
--         is_misinformation boolean,
--         is_misinformation_flagged_at timestamp,
--         created_at timestamp default now (),
--         updated_at timestamp default now ()
--     );
-- create table
--     if not exists public.follows (
--         user_id bigint not null references public.users (id),
--         follower_id bigint not null references public.users (id),
--         created_at timestamp default now (),
--         updated_at timestamp default now (),
--         unique (user_id, follower_id)
--     );
-- create index posts_user_id_index on public.posts (user_id);
-- create index follows_user_id_index on public.follows (user_id);
-- create index follows_follower_id_index on public.follows (follower_id);