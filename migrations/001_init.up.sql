CREATE TABLE IF NOT EXISTS lists
(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,

    item_id INTEGER REFERENCES items (id)
);

CREATE TABLE IF NOT EXISTS items
(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

