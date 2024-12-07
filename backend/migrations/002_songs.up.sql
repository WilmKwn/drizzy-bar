DROP TABLE IF EXISTS songs;

CREATE TABLE IF NOT EXISTS songs (
    title VARCHAR PRIMARY KEY,
    album_name VARCHAR REFERENCES albums(album_name)
);