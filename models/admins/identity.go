package admins

import (
	"github.com/qor/auth/auth_identity"
	"github.com/storebeacon/backend/paola/database"
)

// GetAuthIdentity fetches admin user auth identity by user uID
func GetAuthIdentity(uid string) (*auth_identity.AuthIdentity, error) {
	var authAdmin auth_identity.AuthIdentity
	err := database.Conn.
		Where("uid = ?", uid).
		Find(&authAdmin).
		Error

	return &authAdmin, err
}
