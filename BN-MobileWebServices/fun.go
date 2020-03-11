package main

import (
	"crypto/rand"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

func tokenGenerator() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
func check(newValues string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"contact_info.user_phone": newValues}, person)
	if person.Contacts.UserPhone != "" {
		return false
	}
	connection.Collection("users").FindOne(bson.M{"user_info.user_mail": newValues}, person)
	if person.UserInfos.UserMail != "" {
		return false
	}
	return true
}
func checkBeaconType(beaconType int) string {
	if beaconType == 0 {
		return "Tasma"
	}
	if beaconType == 1 {
		return "Bileklik"
	}
	if beaconType == 5 {
		return "Kalemlik"
	}

	return ""
}
func checkObjID(id string) string {
	var s = bson.IsObjectIdHex(id)
	if s == true {
		return id
	}
	return "123456789123456789123456"
}
