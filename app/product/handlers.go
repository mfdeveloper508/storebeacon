package product

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"github.com/qor/render"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/models/aisle"
	"github.com/storebeacon/backend/models/product"
	"github.com/storebeacon/backend/models/store"
	"github.com/storebeacon/backend/paola/database"
	"github.com/storebeacon/backend/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
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
	elms, err := product.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("product", map[string]interface{}{
		"Products": elms,
		"StoreID": storeId,
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
	elm, err := product.GetRecordByIndex(uint(index))
	if err != nil {
		return
	}

	elmbs, err := aisle.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("edit", map[string]interface{}{
		"Product": elm,
		"Aisles": elmbs,
		"StoreID": storeId,
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
	elm, err := product.GetRecordByIndex(uint(index))
	if err != nil {
		return
	}
	err = database.Conn.Delete(&elm).Error
	if err != nil {
		//requestErrorResponse(w, err.Error())
		return
	}

	// Find the store
	elms, err := product.GetRecordsByStore(uint(storeId))
	if err != nil {
		return
	}

	ctrl.View.Execute("product", map[string]interface{}{
		"Products": elms,
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
	name := r.FormValue("name")
	price := r.FormValue("price")
	description := r.FormValue("description")
	aisleId := r.FormValue("aisleId")
	id := r.FormValue("id")
	storeId := r.FormValue("storeId")
	imagename := r.FormValue("imagename")

	if id == "0" {
		var elm product.Product
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

		elm.Name = name
		i, _ = strconv.Atoi(price)
		elm.Price = float32(i)
		elm.Description = description

		err = database.Conn.Create(&elm).Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}

		if imagename != "" {
			elm.File.FileName = imagename
			elm.File.Url = "https://s3." + aws_info.AwsRegion + ".amazonaws.com/" + aws_info.AwsBucket + "/" + imagename
		}
		err = database.Conn.Save(&elm).Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}

	} else {
		i, err := strconv.Atoi(id)
		elm, err := product.GetRecordByIndex(uint(i))
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

		elm.Name = name
		i, _ = strconv.Atoi(price)
		elm.Price = float32(i)
		elm.Description = description

		if imagename != "" {
			elm.File.FileName = imagename
			elm.File.Url = "https://s3." + aws_info.AwsRegion + ".amazonaws.com/" + aws_info.AwsBucket + "/" + imagename
		}
		err = database.Conn.Save(&elm).Error
		if err != nil {
			//requestErrorResponse(w, err.Error())
			return
		}
	}

	http.Redirect(w, r, "/products/", http.StatusSeeOther)

}


// UploadPOST
func UploadPOST(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	id := r.FormValue("id")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		//saveS3(w, file, handle, id+"/"+handle.Filename)
		saveS3(w, file, handle, id+"/"+handle.Filename)
	case "image/png":
		saveS3(w, file, handle, id+"/"+handle.Filename)
	default:
		jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
	}
}

func CVSUploadPOST(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	id := r.FormValue("id")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	saveFile(w, file, handle, id+"/"+handle.Filename)
}

type CVSProduct struct {
	name		string
	price  		float32
	description	string
	aisleid		int
	storeid		int
	imageNmae   string
}

func BulkUploadPOST(w http.ResponseWriter, r *http.Request) {
	storeId := r.FormValue("storeId")
	cvsfile := r.FormValue("cvsfile")
	count := 0

	cvsFilePath := "./public/system/products/"+cvsfile
	if _, err := os.Stat(cvsFilePath); err == nil {
		csvFile, err := os.Open(cvsFilePath)
		if err != nil {
			fmt.Println(err)
		}
		defer csvFile.Close()

		reader := csv.NewReader(csvFile)
		reader.FieldsPerRecord = -1

		csvData, err := reader.ReadAll()
		if err != nil {
			fmt.Println(err)
		}

		var elm CVSProduct
		var elms []CVSProduct

		for i, each := range csvData {
			if i == 0 {
				continue /* header */
			}

			elm.name = each[0]
			_price, _ := strconv.ParseFloat(each[1], 32)
			elm.price = float32(_price)
			elm.description = each[2]
			elm.aisleid, _ = strconv.Atoi(each[3])
			elm.imageNmae = each[4]
			elm.storeid, _ = strconv.Atoi(storeId)
			elms = append(elms, elm)
		}

		for _, elm := range elms {
			var pro product.Product
			pro.AisleID = uint(elm.aisleid)
			err = database.Conn.
				Where("id = ?", pro.AisleID).
				Preload("Store").
				Find(&pro.Aisle).
				Error
			if err != nil || pro.Aisle.StoreID != uint(elm.storeid) {
				continue
			}
			pro.Description = elm.description
			pro.File.FileName = elm.imageNmae
			pro.File.Url = "https://s3." + aws_info.AwsRegion + ".amazonaws.com/" + aws_info.AwsBucket + "/" + elm.imageNmae

			pro.Name = elm.name
			pro.Price = elm.price
			pro.StoreID = uint(elm.storeid)
			err = database.Conn.
				Where("id = ?", pro.StoreID).
				Find(&pro.Store).
				Error
			if err != nil{
				continue
			}
			err = database.Conn.Save(&pro).Error

			if err == nil {
				count++
			}
		}
	}
	if count == 0 {
		jsonResponse(w, http.StatusOK, "Please check the csv file. No Product uploaded.")
	} else {
		jsonResponse(w, http.StatusOK, strconv.Itoa(count) + " Products uploaded successfully!")
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader, filename string) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = ioutil.WriteFile("./public/system/products/"+filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	jsonResponse(w, http.StatusOK, "File uploaded successfully!.")
}

func saveS3(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader, filename string) {
	token := ""
	creds := credentials.NewStaticCredentials(aws_info.AwsAccess, aws_info.AwsSecret, token)
	_, err := creds.Get()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	cfg := aws.NewConfig().WithRegion(aws_info.AwsRegion).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	var size int64 = handle.Size
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	path := handle.Filename
	params := &s3.PutObjectInput{
		ACL: aws.String(aws_info.AwsRule),
		Bucket: aws.String(aws_info.AwsBucket),
		Key: aws.String(path),
		Body: fileBytes,
		ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Printf("response : " + awsutil.StringValue(resp))


	jsonResponse(w, http.StatusOK, "File uploaded successfully!.")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}