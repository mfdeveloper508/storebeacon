package aisle

import (
	"github.com/jinzhu/gorm"
	"github.com/storebeacon/backend/models/store"
	"github.com/storebeacon/backend/paola/database"
)

// Aisle model
type Aisle struct {
	gorm.Model
	Name    		string `gorm:"size:255" json:"name" valid:"required"`
	Description    	string
	StoreID 		uint
	Store			store.Store
}

func GetRecordByIndex(id uint) (*Aisle, error) {
	var record Aisle
	err := database.Conn.
		Where("id = ?", id).
		Preload("Store").
		Find(&record).
		Error

	return &record, err
}

func GetRecordsByStore(store uint) ([]*Aisle, error) {
	aisles := make([]*Aisle, 0)
	err := database.Conn.
		Where("store_id = ?", store).
		Preload("Store").
		Find(&aisles).
		Error

	return aisles, err
}