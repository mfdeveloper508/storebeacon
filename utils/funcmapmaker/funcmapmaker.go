package funcmapmaker

import (
	"html/template"
	"net/http"

	"github.com/qor/render"
	"github.com/storebeacon/backend/utils"
	"strings"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/models/store"
	"os"
	"github.com/storebeacon/backend/paola/database"
)

// AddFuncMapMaker add FuncMapMaker to view
func AddFuncMapMaker(view *render.Render) *render.Render {
	oldFuncMapMaker := view.FuncMapMaker

	view.FuncMapMaker = func(render *render.Render, r *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}
		if oldFuncMapMaker != nil {
			funcMap = oldFuncMapMaker(render, r, w)
		}

		funcMap["raw"] = func(str string) template.HTML {
			return template.HTML(utils.HTMLSanitizer.Sanitize(str))
		}

		funcMap["get_frontend"] = func() bool {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[1], "auth") || strings.Contains(result[1], "verify"){
					return false;
				}
			}
			return true
		}

		funcMap["get_aisle"] = func() bool {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[1], "aisle"){
					return true;
				}
			}
			return false
		}

		funcMap["get_stores"] = func() bool {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[1], "store"){
					return true;
				}
			}
			return false
		}

		funcMap["get_products"] = func() bool {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[1], "products"){
					return true;
				}
			}
			return false
		}

		funcMap["get_beacons"] = func() bool {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[1], "beacons"){
					return true;
				}
			}
			return false
		}

		funcMap["get_store"] = func() store.Store{
			token := utils.GetToken(r);
			storeId := api.ForceUILoginMiddleware(w, r, token);

			var record store.Store
			err := database.Conn.
				Where("id = ?", storeId).
				Find(&record).
				Error

			elm, err := store.GetRecordByIndex(uint(storeId))
			if err!= nil {
				return record
			}

			if _, err := os.Stat(elm.File.Url); err == nil {
				record.File.Url = ""
			}

			return record
		}

		return funcMap
	}

	return view
}
