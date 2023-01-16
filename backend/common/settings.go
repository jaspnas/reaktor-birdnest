package common

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Settings struct {
	Port            int    `json:"port"`
	DbUrl           string `json:"db_url"`
	DbUname         string `json:"db_uname"`
	DbPass          string `json:"db_pass"`
	DbName          string `json:"db_name"`
	DroneCollection string `json:"drone_collection"`
	GridSize        int    `json:"grid_size"`
	BirdPos         []int  `json:"bird_pos"`
	NoFly           int    `json:"no_fly"`
}

func ReadSettings() Settings {
	jsonFile, err := os.Open("settings.json")
	if err != nil {
		log.Println(err)
	}
	byteData, _ := io.ReadAll(jsonFile)
	err = jsonFile.Close()
	if err != nil {
		log.Println(err)
	}
	var settings Settings
	err = json.Unmarshal(byteData, &settings)
	if err != nil {
		// Return default settings object
		return Settings{
			Port:            8080,
			DbUrl:           "mongodb://localhost:27017",
			DbUname:         "test",
			DbPass:          "Example_Pass123",
			DbName:          "test",
			DroneCollection: "drones",
			GridSize:        500000,
			BirdPos:         []int{250000, 250000},
			NoFly:           100000,
		}
	}
	return settings
}
