CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid NOT NULL UNIQUE PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  elo SMALLINT NOT NULL,

  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

INSERT INTO users (id, username, email, password, elo, created_at, updated_at) 
VALUES ('0dde213a-a81b-4b02-a665-41ec6c037112', 'Krosantos', 'tymko.tecca@gmail.com', 'hashyhashyhashy', 1500, NOW(), NOW());

INSERT INTO users (id, username, email, password, elo, created_at, updated_at) 
VALUES (uuid_generate_v4(), 'Ramira', 'rgutierrez@gmail.com', 'hashyhashyhashy', 1500, NOW(), NOW());

INSERT INTO users (id, username, email, password, elo, created_at, updated_at) 
VALUES (uuid_generate_v4(), 'Castillian', 'lemon.cake@hotmail.com', 'hashyhashyhashy', 1500, NOW(), NOW());

INSERT INTO users (id, username, email, password, elo, created_at, updated_at) 
VALUES (uuid_generate_v4(), 'Hjeregard', 'kvark@jarp.sv', 'hashyhashyhashy', 1500, NOW(), NOW());
