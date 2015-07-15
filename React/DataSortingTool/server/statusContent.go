package main

import(
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"
)

type StatusObject struct {
	Data []struct {
		From struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		ID    string `json:"id"`
		Likes struct {
			Data []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					After  string `json:"after"`
					Before string `json:"before"`
				} `json:"cursors"`
			} `json:"paging"`
		} `json:"likes"`
		Message string `json:"message"`
		Tags    struct {
			Data []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					After  string `json:"after"`
					Before string `json:"before"`
				} `json:"cursors"`
			} `json:"paging"`
		} `json:"tags"`
		UpdatedTime string `json:"updated_time"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			After  string `json:"after"`
			Before string `json:"before"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

var statusData *StatusObject

func statusJSONToStruct(response *http.Response){
     fmt.Println("Handling StatusJSONToStruct")

     responseObject, err := ioutil.ReadAll(response.Body)
    handleError(err, "statusContent.statusJSONToStruct")

     statusData = new (StatusObject)

     err = json.Unmarshal([]byte(responseObject), &statusData)
    handleError(err, "statusContent.statusJSONToStruct")
}

