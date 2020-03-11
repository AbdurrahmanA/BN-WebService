package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func addlostbeaconspage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var devices = getMyDevicesandInfos(r.FormValue("id"))
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
func getMyDevicesandInfos(getID string) []byte {
	var data []byte
	var l []*MyDevicesDetailAndInfos
	var user *MyDevicesDetailAndInfos
	conroltID := checkObjID(getID)
	id := bson.ObjectIdHex(conroltID)
	beacon := &Beacon{}
	beacons := connection.Collection("beacons").Find(bson.M{"user.user_id": id})
	for beacons.Next(beacon) {
		user = &MyDevicesDetailAndInfos{beacon.Id, beacon.Information.BeaconName, beacon.UserInfos.UserMail, beacon.UserInfos.UserPhone}
		l = append(l, user)
	}
	if l != nil {
		data, _ = json.Marshal(l)
	}
	fmt.Println(string(data))
	return data
}
