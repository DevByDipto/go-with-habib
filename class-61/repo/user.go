package repo

import "errors"

// User represents the data model for a user
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

// UserRepo defines the interface for user data operations
type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}

// userRepo is the concrete implementation (in-memory)
type userRepo struct {
	users []User
}

// NewUserRepo initializes a new repository
func NewUserRepo() UserRepo {
	return &userRepo{}
}

// Create adds a new user to the in-memory slice
func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}

// Find searches for a user by email and password
func (r *userRepo) Find(email, pass string) (*User, error) {
	for i := range r.users {
		if r.users[i].Email == email && r.users[i].Password == pass {
			return &r.users[i], nil
		}
	}
	return nil, errors.New("user not found")
}