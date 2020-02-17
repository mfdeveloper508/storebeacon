package aisle

import (
	"github.com/qor/render"
	"net/http"
	"github.com/storebeacon/backend/models/aisle"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/storebeacon/backend/paola/database"
	"github.com/storebeacon/backend/utils"
	"github.com/storebeacon/backend/app/api"
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

	elms, err := aisle.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("aisle", map[string]interface{}{
		"Aisles": elms,
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
	elm, err := aisle.GetRecordByIndex(uint(index))
	if err != nil {
		return
	}
	err = database.Conn.Delete(&elm).Error
	if err != nil {
		//requestErrorResponse(w, err.Error())
		return
	}

	elms, err := aisle.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("aisle", map[string]interface{}{
		"Aisles": elms,
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
	elm, err := aisle.GetRecordByIndex(uint(index))
	if err != nil {
		return
	}

	ctrl.View.Execute("edit", map[string]interface{}{
		"Aisle": elm,
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

	ctrl.View.Execute("edit", map[string]interface{}{
		"StoreID": storeId,
	}, r, w)
}

func (ctrl Controller) Register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	aisleId := r.FormValue("aisleId")
	storeId := r.FormValue("storeId")
	description := r.FormValue("description")

	if aisleId == "0" {
		var elm aisle.Aisle
		i, err := strconv.Atoi(storeId)
		if err != nil {
			http.Redirect(w, r, "/aisle/?error=paramters error", http.StatusSeeOther)
			return
		}
		elm.StoreID = uint(i)
		elm.Description = description
		elm.Name = name

		err = database.Conn.Create(&elm).Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}
	} else {
		i, err := strconv.Atoi(aisleId)
		elm, err := aisle.GetRecordByIndex(uint(i))
		if err != nil {
			http.Redirect(w, r, "/aisle/?error=paramters error", http.StatusSeeOther)
			return
		}

		elm.Description = description
		elm.Name = name

		err = database.Conn.Save(&elm).Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}
	}

	http.Redirect(w, r, "/aisle/", http.StatusSeeOther)

}