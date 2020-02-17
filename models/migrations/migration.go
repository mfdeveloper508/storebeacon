package migrations

import (
	"github.com/qor/auth/auth_identity"

	"github.com/jinzhu/gorm"
	"github.com/storebeacon/backend/models/admins"
	"github.com/storebeacon/backend/models/store"
	"github.com/storebeacon/backend/models/aisle"
	"github.com/storebeacon/backend/models/beacon"
	"github.com/storebeacon/backend/models/product"
	"github.com/storebeacon/backend/paola/database"
)

func Migrate(db *gorm.DB) {
	autoMigrate(&auth_identity.AuthIdentity{})

	autoMigrate(&admins.Admin{})
	autoMigrate(&store.Store{})
	autoMigrate(&aisle.Aisle{})
	autoMigrate(&beacon.Beacon{})
	autoMigrate(&product.Product{})
}

// autoMigrate runs automigrate on provided objects
func autoMigrate(values ...interface{}) {
	for _, value := range values {
		database.Conn.AutoMigrate(value)
	}
}
