package user

const CreateUserTableSQL = `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR,
	last_name VARCHAR, 
	username VARCHAR,
	email VARCHAR,
	password VARCHAR,
	refresh_token VARCHAR
	);`

const DropUserTableSQL = `
	DROP TABLE users;
	`
