package user

const CreateUserTableSQL = `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR NOT NULL,
	last_name VARCHAR NOT NULL, 
	username VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
	refresh_token VARCHAR
	);`

const DropUserTableSQL = `
	DROP TABLE users;
	`
