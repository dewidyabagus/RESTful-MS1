package user

type Service interface {
	// Find user by user id
	FindUserByUserId(id *string) (*User, error)

	// Get user with email and password. for user login
	GetUserWithEmailPassword(email *string, password *string) (id *string, err error)

	// Add new user
	AddNewUser(user *UserAddSpec) error

	// Update user by id
	// UpdateUser(id *string, user *User) error

	// Delete user by id
	// DeleteUser(id *string) error
}

type Repository interface {
	// Find user by user id
	FindUserByUserId(id *string) (*User, error)

	// Get user with email and password. for user login
	GetUserWithEmailPassword(email *string, password *string) (id *string, err error)

	// Add new user
	AddNewUser(user *User) error

	// Update user by id
	// UpdateUser(id *string, user *User) error

	// Delete user by id
	// DeleteUser(id *string) error
}
