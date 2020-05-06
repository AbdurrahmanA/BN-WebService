package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//Message ss
type Message struct {
	AppID            string `json:"app_id"`
	IncludedSegments string `json:"included_segments"`
	Contents         struct {
		En string `json:"en"`
	} `json:"contents"`
}

//AppID ss
var AppID = "56914a6b-2697-4b79-a35c-0ecc952047c3"

func (msg Message) pushNotificationAllUsers(msgForUsers string) {
	client := http.Client{}
	msg.AppID = AppID
	msg.IncludedSegments = "All"
	msg.Contents.En = msgForUsers
	out, _ := json.Marshal(msg)
	reqBody := strings.NewReader(string(out))
	request, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", reqBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic NDBjM2I0YTMtNDNkNS00NTgwLWE2MWYtOGNkY2MxNzUyYTdk")
	if err != nil {
		log.Println(err.Error())
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
}
