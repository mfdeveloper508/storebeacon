package store

import (
	"github.com/qor/render"
	"net/http"
	"github.com/storebeacon/backend/models/store"
	"github.com/storebeacon/backend/paola/database"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/utils"
	"fmt"
	"strconv"
	"mime/multipart"
	"io/ioutil"
	"os"
)

// Controller news controller
type Controller struct {
	View *render.Render
}

func (ctrl Controller) Home(w http.ResponseWriter, r *http.Request) {

	token := utils.GetToken(r);
	storeId := api.ForceUILoginMiddleware(w, r, token);
	if storeId == 0 {
		http.Redirect(w, r, "/verify/login/", http.StatusSeeOther)
		return
	}

	// Find the store
	elm, err := store.GetRecordByIndex(uint(storeId))
	if err != nil || elm.ID == 0{
		return
	}

	ctrl.View.Execute("home", map[string]interface{}{
	}, r, w)
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
	elm, err := store.GetRecordByIndex(uint(storeId))
	if err != nil || elm.ID == 0{
		return
	}

	ctrl.View.Execute("store", map[string]interface{}{
		"Store": elm,
	}, r, w)
}

// Store page
func (ctrl Controller) Login(w http.ResponseWriter, r *http.Request) {

	ctrl.View.Execute("login", map[string]interface{}{
	}, r, w)
}

// Store page
func (ctrl Controller) Store(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	var elm store.Store
	if err := database.Conn.Where("email = ?", email).Find(&elm).Error; err != nil {
		ctrl.View.Execute("login", map[string]interface{}{
			"AuthError": "Email incorrect",
		}, r, w)
		return
	}

	if elm.Password != password {
		ctrl.View.Execute("login", map[string]interface{}{
			"AuthError": "Password incorrect",
		}, r, w)
		return
	}

	token, err := api.GenerateToken(elm.ID)
	if err != nil {
		ctrl.View.Execute("login", map[string]interface{}{
			"AuthError": "Token incorrect",
		}, r, w)
		return
	}


	utils.AddToken(w, r, token);
	http.Redirect(w, r, "/store/", http.StatusSeeOther)
	return
}

// RegisterGET creates a user
func (ctrl Controller) Register(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		http.Redirect(w, r, "/store/?error=paramters error", http.StatusSeeOther)
		return
	}

	elm, err := store.GetRecordByIndex(uint(index))
	if err != nil || elm.Email == ""{
		http.Redirect(w, r, "/store/?error=paramters error", http.StatusSeeOther)
		return
	}

	elm.Title = r.FormValue("title")
	elm.Email = r.FormValue("email")
	elm.Address = r.FormValue("address")
	elm.Password = r.FormValue("password")
	elm.Description = r.FormValue("description")
	elm.Latitude = r.FormValue("latitude")
	elm.Longitude = r.FormValue("longitude")
	if _, err := os.Stat("./public/system/stores/" + id + ".jpg"); err == nil {
		elm.File.Url = "/system/stores/" + id + ".jpg"
		elm.File.FileName = id + ".jpg"
	}
	if _, err := os.Stat("./public/system/stores/" + id + ".png"); err == nil {
		elm.File.Url = "/system/stores/" + id + ".png"
		elm.File.FileName = id + ".png"
	}

	err = database.Conn.Save(&elm).Error
	if err != nil {
		//requestErrorResponse(w, err.Error())
		return
	}

	http.Redirect(w, r, "/store/", http.StatusSeeOther)
}

// UploadPOST
func UploadPOST(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	id := r.FormValue("id")
	fmt.Println(id)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		saveFile(w, file, handle, id+".jpg")
	case "image/png":
		saveFile(w, file, handle, id+".png")
	default:
		jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader, filename string) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = ioutil.WriteFile("./public/system/stores/"+filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}