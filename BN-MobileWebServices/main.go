package main

import (
	"log"
	"net/http"
)

func main() {

	routers.HandleFunc("/login", login).Methods("POST")
	routers.HandleFunc("/update", updateInfo).Methods("GET")
	routers.HandleFunc("/updatedevice", updateDeviceInfo).Methods("GET")
	routers.HandleFunc("/profile", profileInfos).Methods("POST")
	routers.HandleFunc("/devices", myDevices).Methods("POST")
	routers.HandleFunc("/devicesdetail", myDevicesDetails).Methods("POST")
	routers.HandleFunc("/lostpage", lostDeviceList)
	routers.HandleFunc("/addlostbeacon", addlostbeaconspage)

	log.Fatal(http.ListenAndServe(":8090", routers))

}
