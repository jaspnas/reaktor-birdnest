package drone_updater

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Pilot struct {
	PilotId     string `json:"pilotId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	CreatedAt   string `json:"createdDt"`
	Email       string `json:"email"`
}

func getPilot(drone Drone) Pilot {

	serial := drone.SerialNumber

	res, err := http.Get("https://assignments.reaktor.com/birdnest/pilots/" + serial)
	if err != nil {
		log.Println(err)
	}

	var pilot Pilot
	data, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(data, &pilot)
	if err != nil {
		log.Println(err)
	}

	return pilot

}
