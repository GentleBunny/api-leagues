CREATE DATABASE ultra_manager
ENCODING = 'UTF8';

\connect ultra_manager

CREATE TABLE app_user (
	id SERIAL PRIMARY KEY,
	name VARCHAR(45),
	username VARCHAR(45),
	mail VARCHAR(255),
	age SMALLINT,
	gender CHAR(1)
);

INSERT INTO app_user (name, username, mail, age, gender) VALUES ('test', 'test', 'test@mail.com', 26, 'M');