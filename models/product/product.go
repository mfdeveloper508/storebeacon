package product

import (
	"github.com/jinzhu/gorm"
	"github.com/storebeacon/backend/models/aisle"
	"github.com/storebeacon/backend/paola/database"
	"github.com/qor/media/media_library"
	"github.com/storebeacon/backend/models/store"
)

// Product model
type Product struct {
	gorm.Model
	Name		string		`gorm:"size:64" json:"name"`
	Price     	float32    	`gorm:"size:255;index" json:"price"`
	Description	string		`gorm:"size:294967295" json:"description"`
	File  media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}.{{extension}}"`
	AisleID 	uint
	Aisle		aisle.Aisle
	StoreID 	uint
	Store		store.Store
}

func GetRecordByIndex(id uint) (*Product, error) {
	var record Product
	err := database.Conn.
		Where("id = ?", id).
		Preload("Store").
		Preload("Aisle").
		Find(&record).
		Error

	return &record, err
}

// BeforeUpdate hashes the password on update
func (u *Product) BeforeUpdate(scope *gorm.Scope) error {

	if u.File.FileName == "" {
		var store Product
		database.Conn.Where("id = ?", u.ID).Find(&store)
		u.File = store.File
	}

	return nil
}

func GetRecordsByStore(store uint) ([]*Product, error) {
	beacons := make([]*Product, 0)
	err := database.Conn.
		Where("store_id = ?", store).
		Preload("Store").
		Preload("Aisle").
		Find(&beacons).
		Error

	return beacons, err
}