CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         firstname VARCHAR(64),
                         lastname VARCHAR(64),
                         age INTEGER
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(256),
                       genre VARCHAR(64),
                       description TEXT,
                       release_year INTEGER,
                       author_id INTEGER,
                       FOREIGN KEY(author_id) REFERENCES authors(id)
)