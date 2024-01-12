CREATE TABLE IF NOT EXISTS rooms (
   id serial PRIMARY KEY,
   code VARCHAR(25) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
    id serial PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    price FLOAT NOT NULL,
    quantity INTEGER NOT NULL,
    room_id INTEGER REFERENCES rooms(id) ON DELETE CASCADE NOT NULL
);