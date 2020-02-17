package admin

import (
	"github.com/qor/admin"
	"github.com/storebeacon/backend/models/admins"
)

// SetupAdminUsers setup worker
func SetupAdminUsers(Admin *admin.Admin) {
	adminUser := Admin.AddResource(&admins.Admin{}, &admin.Config{Menu: []string{"Site Management"}})

	// Update input types
	adminUser.Meta(&admin.Meta{Name: "Password", Type: "password"})

	adminUser.Meta(&admin.Meta{
		Name: "Role",
		Type: "select_one",
		Config: &admin.SelectOneConfig{
			Collection: []string{"Admin"},
		},
	})

	// Structure the new form
	adminUser.NewAttrs("ID", "Name", "Email", "Password", "Role")

	// Structure the edit form
	adminUser.EditAttrs("ID", "Name", "Email", "Password", "Role")

	// Show given attributes
	adminUser.IndexAttrs("ID", "Name", "Email", "Role")
}
