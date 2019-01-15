package stores

import (
	"database/sql"
	"github.com/avi1989/Identity/models"
)

type Store struct {
	getDb func() *sql.DB
}

type IUserStore interface {
	GetUser(id int) (models.User, error)
	GetUserByEmailAddress(emailAddress string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	AddUser(user *models.User) (int, error)
}

func (store *Store) GetUser(id int) (models.User, error) {
	var user models.User
	database := store.getDb()
	err := database.QueryRow(
		`SELECT id, email_address, password, first_name, last_name, is_active, is_email_verified, role_code
				FROM identity.user where id = $1`, id).Scan(&user.Id, &user.EmailAddress, &user.Password, &user.FirstName, &user.LastName, &user.IsActive, &user.IsEmailVerified, &user.RoleCode)
	if err != nil {
		return user, err
	}
	defer database.Close()
	return user, nil
}

func (store *Store) GetUserByEmailAddress(emailAddress string) (models.User, error) {
	var user models.User
	database := store.getDb()
	err := database.QueryRow(
		`SELECT id, email_address, password, first_name, last_name, is_active, is_email_verified, role_code
				FROM identity.user where email_address = $1`, emailAddress).Scan(&user.Id, &user.EmailAddress, &user.Password, &user.FirstName, &user.LastName, &user.IsActive, &user.IsEmailVerified, &user.RoleCode)
	if err != nil {
		return user, err
	}
	defer database.Close()
	return user, nil
}

func (store *Store) GetAllUsers() ([]models.User, error) {
	db := store.getDb()
	var users []models.User
	rows, err := db.Query(
		`SELECT id, email_address, password, first_name, last_name, is_active, is_email_verified, role_code FROM identity.user `)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.EmailAddress, &user.Password, &user.FirstName, &user.LastName, &user.IsActive, &user.IsEmailVerified, &user.RoleCode)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	defer db.Close()
	return users, nil
}

func (store *Store) AddUser(user *models.User) (int, error) {
	db := store.getDb()
	var id int
	rows := db.QueryRow(
		`INSERT INTO identity.user (email_address, password, first_name, last_name, is_active, is_email_verified, role_code) VALUES
						($1, $2, $3, $4, $5, $6, $7)
						RETURNING id`, user.EmailAddress, user.Password, user.FirstName, user.LastName, user.IsActive, user.IsEmailVerified, user.RoleCode)
	err := rows.Scan(&id)
	defer db.Close()
	return id, err
}

func NewUserStore(getDb func() *sql.DB) IUserStore {
	return &Store{getDb: getDb}
}
