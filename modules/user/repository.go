package user

import (
	"RESTfulMS1/business"
	"RESTfulMS1/business/user"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"id;type:uuid;primaryKey"`
	Email     string    `gorm:"email;type:varchar(50);index:users_email_uniq;unique;not null"`
	FirstName string    `gorm:"first_name;type:varchar(100);not null"`
	LastName  string    `gorm:"last_name;type:varchar(100);not null"`
	Password  string    `gorm:"password;type:varchar(32);not null"`
	Role      string    `gorm:"role;type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"created_at;type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"updated_at;type:timestamp;not null"`
	DeletedAt time.Time `gorm:"deleted_at;type:timestamp"`
}

type UserLogin struct {
	ID string `gorm:"id"`
}

func toUserInsert(u *user.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) AddNewUser(user *user.User) error {
	testUser := new(User)

	err := r.DB.First(testUser, "email = ?", user.Email).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			return err
		}
		return business.ErrDataConflict
	}

	return r.DB.Create(toUserInsert(user)).Error
}

func (r *Repository) GetUserWithEmailPassword(email *string, password *string) (id *string, err error) {
	var login = new(UserLogin)

	err = r.DB.Model(&User{}).First(login, "email = ? and password = md5(?)", email, password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return id, business.ErrUnauthorized
		}
		return id, err
	}

	return &login.ID, nil
}
