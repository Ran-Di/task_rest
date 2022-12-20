SET TIME ZONE 'Europe/Moscow';

CREATE TYPE Crypt AS ENUM ('encrypt', 'decrypt');

CREATE TABLE history (
    id SERIAL PRIMARY KEY,
    type Crypt NOT NULL,
    input VARCHAR(80),
    output VARCHAR(80)
);
