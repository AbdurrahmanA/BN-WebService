package main

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func updateInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var update = controlInfo(r.FormValue("newName"), r.FormValue("newSurName"), r.FormValue("newMail"), r.FormValue("newPhone"), r.FormValue("newPass"), r.FormValue("id"))
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

func controlInfo(name string, surname string, mail string, phone string, password string, id string) bool {
	var checkmail = check(mail)
	var checkphone = check(phone)

	if checkmail == true && checkphone == true {
		person := &Person{}
		conroltID := checkObjID(id)
		err := connection.Collection("users").FindById(bson.ObjectIdHex(conroltID), person)
		if err != nil {
			fmt.Println(err.Error())
		}
		person.Contacts.UserRealName = name
		person.Contacts.UserSurname = surname
		person.UserInfos.UserMail = mail
		person.Contacts.UserPhone = phone
		person.UserInfos.UserPassword = password
		arr := connection.Collection("users").Save(person)
		if arr != nil {
			fmt.Println(err.Error())
		}
		return true
	}
	return false
}
