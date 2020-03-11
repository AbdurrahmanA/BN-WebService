package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func myDevicesDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var devices = getMyDeviceDetails(r.FormValue("id"))
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
func getMyDeviceDetails(getID string) []byte {
	beacon := &Beacon{}
	conroltID := checkObjID(getID)
	var data []byte
	err := connection.Collection("beacons").FindById(bson.ObjectIdHex(conroltID), beacon)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	var beaconType = checkBeaconType(beacon.Information.BeaconType)
	beacons := &MyDevicesDetail{beacon.Information.BeaconName, beaconType, beacon.Information.Variance}
	data, _ = json.Marshal(beacons)
	return data

}
