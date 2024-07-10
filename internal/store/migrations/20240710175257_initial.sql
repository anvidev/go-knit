-- +goose Up
CREATE TABLE "users" (
  id SERIAL PRIMARY KEY,
  name VARCHAR (50) NOT NULL,
  email VARCHAR (50) UNIQUE NOT NULL,
  hashed_password VARCHAR (255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TYPE difficulty as ENUM ('easy', 'medium', 'hard');
CREATE TABLE "projects" (
  id SERIAL PRIMARY KEY,
  title VARCHAR (100) NOT NULL,
  description VARCHAR (255),
  designer VARCHAR (50),
  yarn VARCHAR (50),
  needle_size INT NOT NULL,
  difficulty difficulty NOT NULL,
  public BOOLEAN DEFAULT false,
  user_id INT NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES "users"(id)
);

-- +goose Down
DROP TABLE "users";
DROP TABLE "projects";
DROP TYPE "difficulty";
