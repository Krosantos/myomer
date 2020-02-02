CREATE TABLE users (
  id BIGSERIAL NOT NULL UNIQUE PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  salt VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,

  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

INSERT INTO users (username, email, salt, password, created_at, updated_at) 
VALUES ('Krosantos', 'tymko.tecca@gmail.com', 'saltsaltsalt', 'passpasspass', NOW(), NOW());
