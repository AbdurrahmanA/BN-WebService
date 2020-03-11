package main

import (
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
)

func updateDeviceInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var update = controlDeviceInfo(r.FormValue("newName"), r.FormValue("newType"), r.FormValue("newVariance"), r.FormValue("beaconID"))
		if update == true {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`{"messega": successful}`))
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`{"messega": something went wrong}`))
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"messega": something went wrong}`))
	}
}

func controlDeviceInfo(name string, getType string, variance string, beaconID string) bool {
	device := &Beacon{}
	conroltID := checkObjID(beaconID)
	err := connection.Collection("beacons").FindById(bson.ObjectIdHex(conroltID), device)
	if err != nil {
		return false
	}

	newVariance, _ := strconv.Atoi(variance)
	newType, _ := strconv.Atoi(getType)

	device.Information.BeaconName = name
	device.Information.BeaconType = newType
	device.Information.Variance = newVariance
	errors := connection.Collection("beacons").Save(device)
	if errors != nil {
		return false
	}
	return true

}
