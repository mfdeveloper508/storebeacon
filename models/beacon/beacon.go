package beacon

import (
	"github.com/jinzhu/gorm"
	"github.com/storebeacon/backend/paola/database"
	"github.com/storebeacon/backend/models/aisle"
	"github.com/storebeacon/backend/models/store"
)

// Beacon model
type Beacon struct {
	gorm.Model
	UUID	string		`gorm:"size:64" json:"uuid"`
	Major	int32		`gorm:"size:32" json:"major"`
	Minor	int32		`gorm:"size:32" json:"minor"`
	Color	string		`gorm:"size:64" json:"color"`
	Label	string		`gorm:"size:64" json:"label"`
	Promotion	string
	AisleID 		uint
	Aisle			aisle.Aisle
	StoreID 		uint
	Store			store.Store

	Featured			bool
	UsePromotion		bool
}

func GetRecordByIndex(id uint) (*Beacon, error) {
	var record Beacon
	err := database.Conn.
		Where("id = ?", id).
		Preload("Store").
		Preload("Aisle").
		Find(&record).
		Error

	return &record, err
}

func GetRecordsByStore(store uint) ([]*Beacon, error) {
	beacons := make([]*Beacon, 0)
	err := database.Conn.
		Where("store_id = ?", store).
		Preload("Store").
		Preload("Aisle").
		Find(&beacons).
		Error

	return beacons, err
}

func (u *Beacon) BeforeUpdate(scope *gorm.Scope) error {

	if u.Store.ID == 0 {
		database.Conn.Where("id = ?", u.StoreID).Find(&u.Store)
	}

	return nil
}