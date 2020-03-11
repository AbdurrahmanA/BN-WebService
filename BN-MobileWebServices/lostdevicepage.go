package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func lostDeviceList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data = getLostDeviceList(r.FormValue("id"))
		if data == nil {
			fmt.Println("error:", data)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`{"error": true1 }`))
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(string(data)))
		}

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"errors": true }`))
	}
}
func getLostDeviceList(id string) []byte {
	var data []byte
	var beaconInfo *LostBeaconInApp
	conroltID := checkObjID(id)

	var l []*LostBeaconInApp
	beacon := &LostBeacon{}
	lostBeacon := connection.Collection("lost_beacons").Find(bson.M{"user_id": conroltID})
	for lostBeacon.Next(beacon) {
		beaconInfo = &LostBeaconInApp{beacon.LostDate, beacon.LostLat, beacon.LostLong}
		l = append(l, beaconInfo)
	}
	if l != nil {
		data, _ = json.Marshal(l)
	}
	fmt.Println(string(data))
	return data
}
