package models

type User struct {
	Id              int
	EmailAddress    string
	Password        string
	FirstName       string
	LastName        string
	IsActive        bool
	IsEmailVerified bool
	Permissions     []string
	RoleCode        string
}

type Permission struct {
	Code        string
	Description string
}

/*

CREATE SCHEMA identity;

CREATE TABLE identity.role (
  code varchar(25) PRIMARY KEY,
  description TEXT
);

CREATE TABLE identity.permission (
  code varchar(25) PRIMARY KEY,
  description TEXT
);

CREATE TABLE identity.role_permission (
  role_code varchar(25) REFERENCES identity.role,
  permission_code varchar(255) REFERENCES identity.permission,
  PRIMARY KEY (role_code, permission_code)
);


CREATE TABLE identity.user (
  id SERIAL PRIMARY KEY,
  email_address TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  first_name TEXT,
  last_name TEXT,
  is_active bool NOT NULL,
  is_email_verified bool NOT NULL,
  role_code varchar(25) REFERENCES identity.role
);

CREATE INDEX user_ak1 ON identity.user
  ( email_address, is_active )
  INCLUDE (id, password);

*/
