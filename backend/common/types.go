package common

import "time"

type StoredPilot struct {
	Name        string `bson:"name" json:"name"`
	Email       string `bson:"email" json:"email"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
}

type StoredDrone struct {
	SerialNumber    string      `bson:"serial_number" json:"serial_number"`
	LastSeen        time.Time   `bson:"last_seen" json:"last_seen"`
	ClosestDistance float64     `bson:"closest_distance" json:"closest_distance"`
	Pilot           StoredPilot `bson:"pilot" json:"pilot"`
}
