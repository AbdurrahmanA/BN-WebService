package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user = findDB(r.FormValue("id"), r.FormValue("pass"))
		if user == nil {
			fmt.Println("error:", user)
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

func findDB(userName string, userPass string) []byte {
	person := &Person{}
	var data []byte
	result := connection.Collection("users").FindOne(bson.M{"user_info.user_name": userName, "user_info.user_password": userPass}, person)
	if result != nil {
		fmt.Println(result.Error())
	} else {

		person.UserInfos.UserToken = tokenGenerator()
		connection.Collection("users").Save(person)
		user := &Userjon{person.UserInfos.UserToken, false}
		data, _ := json.Marshal(user)
		return data
	}
	return data
}
