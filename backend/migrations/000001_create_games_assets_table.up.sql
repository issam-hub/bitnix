CREATE TABLE IF NOT EXISTS games (
    id uuid PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    price numeric NOT NULL,
    developer_id uuid NOT NULL,
    release_date timestamp(0) with time zone NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);

DO $$ BEGIN
    CREATE TYPE asset_type AS ENUM ('cover', 'trailer', 'download', 'screenshot');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS assets (
    id uuid PRIMARY KEY,
    game_id uuid NOT NULL REFERENCES games ON DELETE CASCADE,
    type asset_type NOT NULL,
    url text NOT NULL,
    filename text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);