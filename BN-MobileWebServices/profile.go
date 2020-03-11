package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func profileInfos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user = profilInfosGet(r.FormValue("token"))
		if user == nil {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`{"error": true }`))
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(string(user)))
		}

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"error": true }`))
	}
}

func profilInfosGet(token string) []byte {
	var data []byte

	person := &Person{}
	err := connection.Collection("users").FindOne(bson.M{"user_info.user_token": token}, person)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		user := &UserInfoInApp{person.Contacts.UserRealName, person.Contacts.UserSurname, person.Contacts.UserPhone, person.UserInfos.UserPassword, person.UserInfos.UserMail}
		data, _ := json.Marshal(user)
		return data
	}
	return data
}
