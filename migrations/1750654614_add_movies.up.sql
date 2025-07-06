CREATE TABLE IF NOT EXISTS movies(
    id uuid primary key,
    title varchar not null,
    description varchar,
    img_url varchar,
    meta jsonb,
    created_at timestamptz
);

CREATE TABLE IF NOT EXISTS genres(
    id uuid primary key,
    name varchar not null
);

CREATE TABLE IF NOT EXISTS movie_genre(
    movie_id uuid references movies(id),
    genre_id uuid references genres(id)
);

CREATE TABLE IF NOT EXISTS reviews(
    id uuid primary key,
    movie_id uuid references movies(id),
    user_id uuid references users(id),
    description text,
    rating integer not null,
    created_at timestamptz not null,
    unique(movie_id, user_id)
);