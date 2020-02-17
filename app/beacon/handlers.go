package beacon

import (
	"github.com/qor/render"
	"net/http"
	"github.com/storebeacon/backend/models/beacon"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/storebeacon/backend/models/aisle"
	"github.com/storebeacon/backend/paola/database"
	"github.com/storebeacon/backend/models/store"
	"fmt"
	"github.com/storebeacon/backend/utils"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/models/product"
	"os"
	"strings"
)

// Controller news controller
type Controller struct {
	View *render.Render
}

// Stores page
func (ctrl Controller) Index(w http.ResponseWriter, r *http.Request) {
	token := utils.GetToken(r);
	storeId := api.ForceUILoginMiddleware(w, r, token);
	if storeId == 0 {
		http.Redirect(w, r, "/verify/login/", http.StatusSeeOther)
		return
	}

	// Find the store
	elms, err := beacon.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("beacon", map[string]interface{}{
		"Beacons": elms,
	}, r, w)
}

func (ctrl Controller) Remove(w http.ResponseWriter, r *http.Request) {
	token := utils.GetToken(r);
	storeId := api.ForceUILoginMiddleware(w, r, token);
	if storeId == 0 {
		http.Redirect(w, r, "/verify/login/", http.StatusSeeOther)
		return
	}
	_index := mux.Vars(r)["id"]
	index, _ := strconv.Atoi(_index)

	// Find the store
	elm, err := beacon.GetRecordByIndex(uint(index))
	if err != nil {
		return
	}
	err = database.Conn.Delete(&elm).Error
	if err != nil {
		//requestErrorResponse(w, err.Error())
		return
	}

	// Find the store
	elms, err := beacon.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("beacon", map[string]interface{}{
		"Beacons": elms,
	}, r, w)
}

// Stores page
func (ctrl Controller) Edit(w http.ResponseWriter, r *http.Request) {
	token := utils.GetToken(r);
	storeId := api.ForceUILoginMiddleware(w, r, token);
	if storeId == 0 {
		http.Redirect(w, r, "/verify/login/", http.StatusSeeOther)
		return
	}
	_index := mux.Vars(r)["id"]
	index, _ := strconv.Atoi(_index)

	// Find the store
	elm, err := beacon.GetRecordByIndex(uint(index))
	if err != nil {
		return
	}

	elmbs, err := aisle.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("edit", map[string]interface{}{
		"Beacon": elm,
		"Aisles": elmbs,
		"StoreID": storeId,
	}, r, w)
}

// Stores page
func (ctrl Controller) Add(w http.ResponseWriter, r *http.Request) {
	token := utils.GetToken(r);
	storeId := api.ForceUILoginMiddleware(w, r, token);
	if storeId == 0 {
		http.Redirect(w, r, "/verify/login/", http.StatusSeeOther)
		return
	}

	elmbs, err := aisle.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("edit", map[string]interface{}{
		"Aisles": elmbs,
		"StoreID": storeId,
	}, r, w)
}

func (ctrl Controller) Register(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("uuid")
	major := r.FormValue("major")
	minor := r.FormValue("minor")
	color := r.FormValue("color")
	label := r.FormValue("label")
	promotion := r.FormValue("promotion")
	aisleId := r.FormValue("aisleId")
	id := r.FormValue("id")
	storeId := r.FormValue("storeId")

	featured := r.FormValue("featured")
	usepromotion := r.FormValue("usepromotion")

	fmt.Println(featured)
	fmt.Println(usepromotion)

	if id == "0" {
		var elm beacon.Beacon
		i, err := strconv.Atoi(storeId)
		if err != nil {
			http.Redirect(w, r, "/aisle/?error=paramters error", http.StatusSeeOther)
			return
		}
		elm.StoreID = uint(i)
		var selm store.Store
		_ = database.Conn.
			Where("id = ?", elm.StoreID).
			Find(&selm).
			Error
		elm.Store = selm

		i, err = strconv.Atoi(aisleId)
		var aelm aisle.Aisle
		elm.AisleID = uint(i)
		err = database.Conn.
			Where("id = ?", elm.AisleID).
			Preload("Store").
			Find(&aelm).
			Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}
		elm.Aisle = aelm

		elm.UUID = uuid
		i, _ = strconv.Atoi(major)
		elm.Major = int32(i)
		i, _ = strconv.Atoi(minor)
		elm.Minor = int32(i)
		elm.Color = color
		elm.Label = label
		elm.Promotion = promotion

		if featured == "True" {
			elm.Featured = true
		} else {
			elm.Featured = false
		}
		if usepromotion == "True" {
			elm.UsePromotion = true
		} else {
			elm.UsePromotion = false
		}

		err = database.Conn.Create(&elm).Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}
	} else {
		i, err := strconv.Atoi(id)
		elm, err := beacon.GetRecordByIndex(uint(i))
		if err != nil {
			http.Redirect(w, r, "/aisle/?error=paramters error", http.StatusSeeOther)
			return
		}

		i, err = strconv.Atoi(aisleId)
		var aelm aisle.Aisle
		elm.AisleID = uint(i)
		err = database.Conn.
			Where("id = ?", elm.AisleID).
			Preload("Store").
			Find(&aelm).
			Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}
		elm.Aisle = aelm

		elm.UUID = uuid
		i, _ = strconv.Atoi(major)
		elm.Major = int32(i)
		i, _ = strconv.Atoi(minor)
		elm.Minor = int32(i)
		elm.Color = color
		elm.Label = label
		elm.Promotion = promotion

		if featured == "True" {
			elm.Featured = true
		} else {
			elm.Featured = false
		}
		if usepromotion == "True" {
			elm.UsePromotion = true
		} else {
			elm.UsePromotion = false
		}

		err = database.Conn.Save(&elm).Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}
	}

	http.Redirect(w, r, "/beacons/", http.StatusSeeOther)

}

type UUID struct {
	UUID		string	`json:"uuid"`
	Major		uint	`json:"major"`
	Minor		uint	`json:"minor"`
}

type PRODUCT struct {
	Id		uint	`json:"id"`
}

var serverPrefix = ""
func PostProc(w http.ResponseWriter, r *http.Request) {

	payload := struct {
		Action	string	`json:"action"`
		Uuid   	UUID	`json:"beacon"`
		Product	PRODUCT	`json:"product"`
	}{}

	if strings.Contains(r.Proto, "HTTP/") {
		serverPrefix = "http://" + r.Host
	} else {
		serverPrefix = "https://" + r.Host
	}

	err := api.UnmarshalPayload(r.Body, &payload)
	defer r.Body.Close()

	if err != nil {
		api.BadRequest(w)
		return
	}

	if payload.Action == "getWelcome" {
		GetWelcome(w, &payload.Uuid)
		return
	}

	if payload.Action == "getPromotion" {
		GetPromotion(w, &payload.Uuid)
		return
	}

	if payload.Action == "getProductList" {
		GetProductList(w, &payload.Uuid)
		return
	}

	if payload.Action == "getProduct" {
		GetProduct(w, &payload.Product)
		return
	}

	return
}
func GetWelcome(w http.ResponseWriter, payload *UUID) {

	if payload.UUID == ""{
		api.BadRequest(w)
		return
	}

	beacons := make([]*beacon.Beacon, 0)
	err := database.Conn.
		Where("lower(uuid) = ? AND major = ? AND minor= ?", strings.ToLower(payload.UUID), payload.Major, payload.Minor).
		Preload("Store").
		Preload("Aisle").
		Find(&beacons).
		Error
	if err != nil{
		api.BadRequest(w)
		return
	}

	response := struct {
		Message 	string     	`json:"msg"`
		Data   		string     	`json:"data"`
	}{
		Message: "ERROR",
		Data: "",
	}

	if len(beacons) > 0 && beacons[0].Featured == true{
		response.Message = "SUCCESS"
		response.Data = string(beacons[0].Store.Title)
	}

	api.ServeJSON(w, response)
}

func GetPromotion(w http.ResponseWriter, payload *UUID) {

	if payload.UUID == ""{
		api.BadRequest(w)
		return
	}

	beacons := make([]*beacon.Beacon, 0)
	err := database.Conn.
		Where("lower(uuid) = ? AND major = ? AND minor= ?", strings.ToLower(payload.UUID), payload.Major, payload.Minor).
		Preload("Store").
		Preload("Aisle").
		Find(&beacons).
		Error
	if err != nil{
		api.BadRequest(w)
		return
	}

	response := struct {
		Message 	string     	`json:"msg"`
		Data   		string     	`json:"data"`
	}{
		Message: "ERROR",
		Data: "",
	}

	if len(beacons) > 0 && beacons[0].UsePromotion == true{
		response.Message = "SUCCESS"
		response.Data = string(beacons[0].Promotion)
	}

	api.ServeJSON(w, response)
}

type productInfo struct {
	Id			uint	`json:"id"`
	Name		string	`json:"name"`
	Price		float32	`json:"price"`
	Description	string	`json:"description"`
	ImageUrl	string	`json:"image_url"`
	AisleLabel	string	`json:"aisle_label"`
}

func GetProductList(w http.ResponseWriter, payload *UUID) {

	if payload.UUID == ""{
		api.BadRequest(w)
		return
	}

	beacons := make([]*beacon.Beacon, 0)
	err := database.Conn.
		Where("lower(uuid) = ? AND major = ? AND minor= ?", strings.ToLower(payload.UUID), payload.Major, payload.Minor).
		Preload("Store").
		Preload("Aisle").
		Find(&beacons).
		Error
	if err != nil{
		api.BadRequest(w)
		return
	}

	response := struct {
		Message 	string     	`json:"msg"`
		Data     []*productInfo	`json:"data"`
	}{
		Message: "ERROR",
		Data: nil,
	}

	if len(beacons) == 0{
		api.ServeJSON(w, response)
		return
	}

	var aisleId = beacons[0].AisleID
	products := make([]*product.Product, 0)
	err = database.Conn.
		Where("aisle_id = ?", aisleId).
		Preload("Store").
		Preload("Aisle").
		Find(&products).
		Error
	if err != nil{
		api.BadRequest(w)
		return
	}

	if len(products) == 0{
		api.ServeJSON(w, response)
		return
	}

	jsonElms := make([]*productInfo, 0)
	for _, elm := range products{
		imageUrl := ""
		if _, err := os.Stat("./public" + elm.File.Url); err == nil {
			imageUrl = serverPrefix + elm.File.Url
		}
		jsonElm := &productInfo{
			Id: elm.ID,
			Name: elm.Name,
			Price: elm.Price,
			Description: elm.Description,
			ImageUrl: imageUrl,
			AisleLabel: elm.Aisle.Name,
		}
		jsonElms = append(jsonElms, jsonElm)
	}

	response.Message = "SUCCESS"
	response.Data = jsonElms

	api.ServeJSON(w, response)
}

func GetProduct(w http.ResponseWriter, payload *PRODUCT) {

	elm, err := product.GetRecordByIndex(uint(payload.Id))
	if err != nil {
		api.BadRequest(w)
		return
	}

	imageUrl := ""
	if _, err := os.Stat("./public" + elm.File.Url); err == nil {
		imageUrl = serverPrefix + elm.File.Url
	}
	jsonElm := &productInfo{
		Id: elm.ID,
		Name: elm.Name,
		Price: elm.Price,
		Description: elm.Description,
		ImageUrl: imageUrl,
		AisleLabel: elm.Aisle.Name,
	}

	response := struct {
		Message 	string     	`json:"msg"`
		Data	*productInfo	`json:"data"`
	}{
		Message: "SUCCESS",
		Data: jsonElm,
	}

	api.ServeJSON(w, response)
}