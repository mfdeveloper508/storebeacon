package utils

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/microcosm-cc/bluemonday"
	"github.com/qor/qor/utils"

	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/storebeacon/backend/paola/database"
)

// HTMLSanitizer HTML sanitizer
var HTMLSanitizer = bluemonday.UGCPolicy()
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   []byte = securecookie.GenerateRandomKey(32)
	store        = sessions.NewCookieStore(key)
)

// GetDB from DB from request
func GetDB(r *http.Request) *gorm.DB {
	if db := utils.GetDBFromRequest(r); db != nil {
		return db
	}

	return database.Conn
}

func GetProductCount(req *http.Request) int {
	session, _ := store.Get(req, "product-info")

	if val, ok := session.Values["count"]; ok {
		t := val.(int)
		return t
	} else {
		return 0
	}
}

func GetToken(req *http.Request) string{
	session, _ := store.Get(req, "token-info")

	if val, ok := session.Values["token"]; ok {
		t := val.(string)
		return t
	} else {
		return ""
	}
}

func AddToken(w http.ResponseWriter, req *http.Request, Token string) {
	session, _ := store.Get(req, "token-info")
	session.Values["token"] = Token

	session.Save(req, w)
}

func AddProduct(w http.ResponseWriter, req *http.Request, productID int) {
	session, _ := store.Get(req, "product-info")

	if CheckProduct(req, productID) == true {
		return
	}

	count := 0
	if val, ok := session.Values["count"]; ok {
		count = val.(int)
	}
	count++
	session.Values["count"] = count

	productstr := strconv.Itoa(productID)
	session.Values[productstr] = true

	session.Save(req, w)
}

func CheckProduct(req *http.Request, productID int) bool {
	session, _ := store.Get(req, "product-info")

	productstr := strconv.Itoa(productID)
	if _, ok := session.Values[productstr]; ok {
		return true
	} else {
		return false
	}
}
