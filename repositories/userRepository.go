package repositories

import (
	"database/sql"
	"log"

	comp "github.com/caiomarinello/ninjaGateway/components"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Implements the Registrar interface:
// it saves a new 'user' in the database.
func (d *UserRepository) Register(newUser comp.User) error {
	
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error generating hashed password: ", err)
		return err
	}

	query := "INSERT INTO users (email, password_hash, first_name, last_name, user_role) VALUES (?, ?, ?, ?, ?)"
	_, err = d.db.Exec(query, newUser.Email, hashedPasswd, newUser.FirstName, newUser.LastName, newUser.Role)
	if err != nil {
		return err
	}

	return nil
}

// Implements the UserFetcher interface.
func (d *UserRepository) FetchUserByEmail(email string) (*comp.User, error) {
	var user comp.User

	userRow := d.db.QueryRow("SELECT user_id, email, password_hash, first_name, last_name, user_role FROM users WHERE email = ?", email)
	err := userRow.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Role)
	if err != nil {
		log.Println("Error with db query row scan: ", err)
		return nil, err
	}
	return &user, err
}