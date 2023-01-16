package handlers

import (
	"backend-go/common"
	"backend-go/handlers/util"
	"encoding/json"
	"log"
	"net/http"
)

func HandleGetDroneRequest(w http.ResponseWriter, r *http.Request) {

	settings := common.ReadSettings()

	db := common.DbParameters{
		Url:             settings.DbUrl,
		Username:        settings.DbUname,
		Password:        settings.DbPass,
		DbName:          settings.DbName,
		DroneCollection: settings.DroneCollection,
	}

	drones := util.GetAllDrones(db)

	data, err := json.Marshal(drones)
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Request-Method", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Access-Control-Request-Methods, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
