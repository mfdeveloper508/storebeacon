package store

import (
	"github.com/jinzhu/gorm"
	"strings"
	"regexp"
	"strconv"
	"github.com/qor/media/media_library"
	"github.com/storebeacon/backend/paola/database"
)

// Store model
type Store struct {
	gorm.Model
	Title     string		`gorm:"size:64" json:"title"`
	Email     string    	`gorm:"size:255;index" json:"email"`
	Password  string    	`gorm:"size:64" json:"-"`
	Description	string		`gorm:"size:294967295" json:"description"`
	Address	string			`gorm:"size:64" json:"address"`
	Latitude string 		`gorm:"size:64" json:"latitude"`
	Longitude string 		`gorm:"size:64" json:"longitude"`
	File  media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}.{{extension}}"`
}

func GetRecordByIndex(id uint) (*Store, error) {
	var record Store
	err := database.Conn.
		Where("id = ?", id).
		Find(&record).
		Error

	return &record, err
}

func latlngToDecimal(coord string, dir string, lat bool) string {
	decimal := 0.0
	negative := false

	if (lat && strings.ToUpper(dir) == "S") || strings.ToUpper(dir) == "W" {
		negative = true
	}

	r, _ := regexp.Compile("^-?([0-9]*?)([0-9]{2,2}\\.[0-9]*)$")

	result := r.FindStringSubmatch(coord)
	deg, _ := strconv.ParseFloat(result[1], 32) // degrees
	min, _ := strconv.ParseFloat(result[2], 32) // minutes & seconds

	// Calculate
	decimal = deg + (min / 60)

	if negative {
		decimal *= -1
	}

	_decimal := strconv.FormatFloat(decimal, 'g', 'g', 32)
	return _decimal
}

// BeforeUpdate hashes the password on update
func (u *Store) BeforeUpdate(scope *gorm.Scope) error {

	if u.File.FileName == "" {
		var store Store
		database.Conn.Where("id = ?", u.ID).Find(&store)
		u.File = store.File
	}

	return nil
}