package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func myDevices(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var devices = getMyDevices(r.FormValue("id"))
		if devices == nil {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`{"error": true }`))
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(string(devices)))
		}

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"error": true }`))
	}
}

func getMyDevices(getID string) []byte {
	var data []byte
	var l []*MyDevices
	var user *MyDevices
	id := bson.ObjectIdHex(getID)
	beacon := &Beacon{}
	beacons := connection.Collection("beacons").Find(bson.M{"user.user_id": id})
	for beacons.Next(beacon) {
		user = &MyDevices{beacon.Id, beacon.Information.BeaconName, beacon.Information.BeaconType}
		l = append(l, user)
	}
	data, _ = json.Marshal(l)
	fmt.Println(string(data))
	return data
}
