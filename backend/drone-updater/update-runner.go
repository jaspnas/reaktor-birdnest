package drone_updater

import (
	"backend-go/common"
	"time"
)

func RunUpdates(area ProtectedArea, params common.DbParameters) {

	// Connect to MongoDB
	connection := common.ConnectMongo(params)
	database := common.GetDatabase(connection, params.DbName)
	collection := database.Collection(params.DroneCollection)

	// Run this while the server is on.
	for {
		updateDrones(area, collection)
		purgeOldEntries(collection, 10*time.Minute)
		time.Sleep(2 * time.Second)
	}

}
