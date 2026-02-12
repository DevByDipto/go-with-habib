package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

// UserRepo defines the interface for user data operations
type UserRepo interface {
	user.UserRepo
}

// userRepo is the concrete implementation (in-memory)
type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo initializes a new repository
func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db : db,
	}
}

// Create adds a new user to the in-memory slice
func (r *userRepo) Create(user domain.User) (*domain.User, error) {
    query := `
        INSERT INTO users (
            first_name, 
            last_name, 
            email, 
            password, 
            is_shop_owner
        )
        VALUES (
            :first_name, 
            :last_name, 
            :email, 
            :password, 
            :is_shop_owner
        )
        RETURNING id`

    var userID int
    // Execute named query and scan the returning ID
    rows, err := r.db.NamedQuery(query, user)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    if rows.Next() {
        if err := rows.Scan(&userID); err != nil {
            return nil, err
        }
    }

    user.ID = userID
    return &user, nil
}

// Find searches for a user by email and password
func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`

	err := r.db.Get(&user, query, email, pass) // db is *sqlx.DB
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no matching user
		}
		return nil, err
	}

	return &user, nil
}