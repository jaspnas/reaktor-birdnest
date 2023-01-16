package util

import (
	"backend-go/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetAllDrones(parameters common.DbParameters) []common.StoredDrone {
	collection := common.GetDatabase(common.ConnectMongo(parameters), parameters.DbName).Collection(parameters.DroneCollection)

	filter := bson.M{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}

	defer cursor.Close(context.TODO())

	var drones []common.StoredDrone

	for cursor.Next(context.TODO()) {
		var drone common.StoredDrone
		if err := cursor.Decode(&drone); err != nil {
			log.Println(err)
		}
		drones = append(drones, drone)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
	}

	return drones

}
