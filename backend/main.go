package main

import (
	"backend-go/common"
	"backend-go/drone-updater"
	"backend-go/handlers"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	settings := common.ReadSettings()

	protect := drone_updater.ProtectedArea{
		CentreX: settings.BirdPos[0],
		CentreY: settings.BirdPos[1],
		Radius:  settings.NoFly,
	}

	db := common.DbParameters{
		Url:             settings.DbUrl,
		Username:        settings.DbUname,
		Password:        settings.DbPass,
		DbName:          settings.DbName,
		DroneCollection: settings.DroneCollection,
	}

	go drone_updater.RunUpdates(protect, db)

	server := ":" + strconv.Itoa(settings.Port)

	http.HandleFunc("/api/drones", handlers.HandleGetDroneRequest)
	http.HandleFunc("/api/ws", handlers.WsEndpoint)

	fmt.Printf("Starting server on port %d", settings.Port)
	go func() {
		if err := http.ListenAndServe(server, nil); err != nil {
			log.Fatal(err)
		}
	}()
	if err := http.ListenAndServeTLS(":44310", "/cert.crt", "/key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
