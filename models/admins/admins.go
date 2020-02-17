package admins

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qor/auth/auth_identity"
	"github.com/storebeacon/backend/paola/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Admin User model
type Admin struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name"`
	Email    string `gorm:"size:255;index" json:"email" valid:"email,required"`
	Role     string `gorm:"size:64" json:"role"`
	Password string `gorm:"size:64" json:"-"`
}

// GetAdminUsers fetches user by admin role
func GetAdminUsers() ([]*Admin, error) {
	admins := make([]*Admin, 0)
	err := database.Conn.
		Where("role = ?", "Admin").
		Find(&admins).
		Error

	return admins, err
}

// Used for qor admin to get current logged user
func (admin Admin) DisplayName() string {
	return admin.Email
}

// BeforeCreate hashes the password on create
func (u *Admin) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Password", generateHash(u.Password))

	return nil
}

// BeforeCreate hashes the password on create
func (u *Admin) AfterCreate(scope *gorm.Scope) error {
	// Create auth identity
	now := time.Now()

	authIdentity := &auth_identity.AuthIdentity{}
	authIdentity.Provider = "password"
	authIdentity.UID = u.Email
	authIdentity.EncryptedPassword = u.Password
	authIdentity.UserID = fmt.Sprint(u.ID)
	authIdentity.ConfirmedAt = &now

	database.Conn.Create(authIdentity)

	return nil
}

// BeforeUpdate hashes the password on update
func (u *Admin) BeforeUpdate(scope *gorm.Scope) error {
	hashedPassword := generateHash(u.Password)

	scope.SetColumn("Password", hashedPassword)

	// Update auth identity
	authAdmin, _ := GetAuthIdentity(u.Email)
	authAdmin.EncryptedPassword = hashedPassword
	database.Conn.Save(authAdmin)

	return nil
}

func generateHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}
