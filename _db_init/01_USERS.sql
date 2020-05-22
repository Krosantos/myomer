CREATE TABLE users (
  id BIGSERIAL NOT NULL UNIQUE PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  elo SMALLINT NOT NULL,

  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

INSERT INTO users (username, email, password, elo, created_at, updated_at) 
VALUES ('Krosantos', 'tymko.tecca@gmail.com', 'hashyhashyhashy', 1500, NOW(), NOW());

INSERT INTO users (username, email, password, elo, created_at, updated_at) 
VALUES ('Ramira', 'rgutierrez@gmail.com', 'hashyhashyhashy', 1500, NOW(), NOW());

INSERT INTO users (username, email, password, elo, created_at, updated_at) 
VALUES ('Castillian', 'lemon.cake@hotmail.com', 'hashyhashyhashy', 1500, NOW(), NOW());

INSERT INTO users (username, email, password, elo, created_at, updated_at) 
VALUES ('Hjeregard', 'kvark@jarp.sv', 'hashyhashyhashy', 1500, NOW(), NOW());
