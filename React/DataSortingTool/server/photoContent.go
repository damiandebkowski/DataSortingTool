package main

import(
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"
)

type PhotoObject struct {
	Data []struct {
		CreatedTime string `json:"created_time"`
		From        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		Height int    `json:"height"`
		Icon   string `json:"icon"`
		ID     string `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			Source string `json:"source"`
			Width  int    `json:"width"`
		} `json:"images"`
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
		Link        string `json:"link"`
		Picture     string `json:"picture"`
		Source      string `json:"source"`
		UpdatedTime string `json:"updated_time"`
		Width       int    `json:"width"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			After  string `json:"after"`
			Before string `json:"before"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

var photoData *PhotoObject

func photoJSONToStruct(response *http.Response){
     fmt.Println("Handling PhotoJSONToStruct...")

     responseObject, err := ioutil.ReadAll(response.Body)
    handleError(err, "photoContent.photoJSONToStruct")

     photoData = new (PhotoObject)

     err = json.Unmarshal([]byte(responseObject), &photoData)
    handleError(err, "photoContent.photoJSONToStruct")

     //test struct
     fmt.Println("Test Photo ID: " + photoData.Data[0].ID)
}

