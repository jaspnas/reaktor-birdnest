package drone_updater

import (
	"backend-go/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func purgeOldEntries(collection *mongo.Collection, duration time.Duration) {
	filter := bson.M{"last_seen": bson.M{"$lt": time.Now().Add(-duration)}}

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
}

func addDrone(collection *mongo.Collection, drone common.StoredDrone) {
	filter := bson.M{"serial_number": drone.SerialNumber}

	var temp common.StoredDrone

	err := collection.FindOne(context.TODO(), filter).Decode(&temp)
	if err != nil {
		_, err2 := collection.InsertOne(context.TODO(), drone)
		if err2 != nil {
			return
		}
	}

	update := bson.M{"$set": bson.M{
		"closest_distance": bson.M{"$min": []float64{drone.ClosestDistance, temp.ClosestDistance}},
		"last_seen":        drone.LastSeen,
	}}

	collection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetUpsert(true))
}
