package drone_updater

import (
	"backend-go/common"
	"backend-go/handlers"
	"encoding/xml"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"math"
	"net/http"
	"time"
)

type ProtectedArea struct {
	CentreX int
	CentreY int
	Radius  int
}

type Drone struct {
	SerialNumber string  `xml:"serialNumber"`
	Model        string  `xml:"model"`
	Manufacturer string  `xml:"manufacturer"`
	MAC          string  `xml:"mac"`
	IPv4         string  `xml:"ipv4"`
	IPv6         string  `xml:"ipv6"`
	Firmware     string  `xml:"firmware"`
	PositionY    float64 `xml:"positionY"`
	PositionX    float64 `xml:"positionX"`
	Altitude     float64 `xml:"altitude"`
}

type Capture struct {
	SnapshotTimestamp string  `xml:"snapshotTimestamp,attr"`
	Drones            []Drone `xml:"drone"`
}

type DroneResponse struct {
	DeviceInformation struct {
		DeviceID         string `xml:"deviceId,attr"`
		ListenRange      int    `xml:"listenRange"`
		DeviceStarted    string `xml:"deviceStarted"`
		UptimeSeconds    int    `xml:"uptimeSeconds"`
		UpdateIntervalMs int    `xml:"updateIntervalMs"`
	} `xml:"deviceInformation"`
	Capture Capture `xml:"capture"`
}

func getDrones(data []byte) ([]Drone, string) {
	var parsedResponse DroneResponse
	err := xml.Unmarshal(data, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return parsedResponse.Capture.Drones, parsedResponse.Capture.SnapshotTimestamp
}

func checkBounds(x float64, y float64, bounds ProtectedArea) bool {
	// Calculate distance using the euclidean distance formula, then check if it is smaller than the radius of the
	//protected area.
	return float64(bounds.Radius) > math.Sqrt(math.Pow(x-float64(bounds.CentreX), 2)+
		math.Pow(y-float64(bounds.CentreY), 2))

}

func checkDrones(drones []Drone, area ProtectedArea) []Drone {
	var infractors []Drone
	for i := 0; i < len(drones); i++ {
		drone := drones[i]
		if checkBounds(drone.PositionX, drone.PositionY, area) {
			infractors = append(infractors, drone)
		}
	}
	return infractors
}

func createStoredData(drone Drone, pilot Pilot, centrepoint ProtectedArea, timestamp string) common.StoredDrone {

	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		t = time.Now()
	}

	distance := math.Sqrt(math.Pow(drone.PositionX-float64(centrepoint.CentreX), 2) +
		math.Pow(drone.PositionY-float64(centrepoint.CentreY), 2))

	sp := common.StoredPilot{
		Name:        pilot.FirstName + " " + pilot.LastName,
		Email:       pilot.Email,
		PhoneNumber: pilot.PhoneNumber,
	}
	return common.StoredDrone{
		SerialNumber:    drone.SerialNumber,
		LastSeen:        t,
		ClosestDistance: distance,
		Pilot:           sp,
	}
}

func sendToWebSockets(drone common.StoredDrone) {

	for client := range handlers.Clients {
		if err := client.WriteJSON(drone); err != nil {
			log.Println(err)
		}
	}

}

func updateDrones(bounds ProtectedArea, database *mongo.Collection) {
	res, err := http.Get("https://assignments.reaktor.com/birdnest/drones")
	if err != nil {
		log.Println(err)
	}

	data, _ := io.ReadAll(res.Body)

	drones, timestamp := getDrones(data)

	infractors := checkDrones(drones, bounds)

	for _, drone := range infractors {
		sd := createStoredData(drone, getPilot(drone), bounds, timestamp)
		addDrone(database, sd)
		sendToWebSockets(sd)
	}

}
