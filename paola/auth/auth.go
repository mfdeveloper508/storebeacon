package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/qor"
	"github.com/qor/roles"

	"github.com/qor/auth/claims"
	"github.com/qor/auth/providers/password"
	"github.com/qor/redirect_back"
	"github.com/storebeacon/backend/models/admins"
	"github.com/storebeacon/backend/paola/database"
)

var Auth *auth.Auth

func NewAuth() *auth.Auth {
	redirectBack := redirect_back.New(&redirect_back.Config{
		IgnoredPrefixes: []string{"/auth"},
		FallbackPath:    "/admin",
	})

	config := &auth.Config{
		DB:         database.Conn,
		UserModel:  admins.Admin{},
		Redirector: auth.Redirector{redirectBack},
	}

	Auth = auth.New(config)

	Auth.RegisterProvider(password.New(&password.Config{
		RegisterHandler: func(*auth.Context) (*claims.Claims, error) {
			return nil, errors.New("Registration disabled")
		},
		ResetPasswordHandler: func(*auth.Context) error {
			return errors.New("Reset disabled")
		},
		RecoverPasswordHandler: func(*auth.Context) error {
			return errors.New("Recover disabled")
		},
	}))

	return Auth
}

func init() {
	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil && currentUser.(*admins.Admin).Role == "Full Admin" && currentUser.(*admins.Admin).DeletedAt == nil
	})
}

type AdminAuth struct {
}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser := Auth.GetCurrentUser(c.Request)
	if currentUser != nil {
		qorCurrentUser, ok := currentUser.(qor.CurrentUser)
		if !ok {
			fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
		}
		return qorCurrentUser
	}
	return nil
}
