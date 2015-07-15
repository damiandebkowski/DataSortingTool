package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type PostObject struct {
	Data []struct {
		Actions []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"actions"`
		Application struct {
			Category  string `json:"category"`
			ID        string `json:"id"`
			Link      string `json:"link"`
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"application"`
		Caption     string `json:"caption"`
		CreatedTime string `json:"created_time"`
		Description string `json:"description"`
		From        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		Icon      string `json:"icon"`
		ID        string `json:"id"`
		IsExpired bool   `json:"is_expired"`
		IsHidden  bool   `json:"is_hidden"`
		Likes     struct {
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
		Link     string `json:"link"`
		Name     string `json:"name"`
		ObjectID string `json:"object_id"`
		Picture  string `json:"picture"`
		Privacy  struct {
			Allow       string `json:"allow"`
			Deny        string `json:"deny"`
			Description string `json:"description"`
			Friends     string `json:"friends"`
			Value       string `json:"value"`
		} `json:"privacy"`
		StatusType string `json:"status_type"`
		Story      string `json:"story"`
		StoryTags  struct {
			_ []struct {
				ID     string `json:"id"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Offset int    `json:"offset"`
				Type   string `json:"type"`
			} `json:"0"`
		} `json:"story_tags"`
		Subscribed  bool   `json:"subscribed"`
		Type        string `json:"type"`
		UpdatedTime string `json:"updated_time"`
	} `json:"data"`
	Paging struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"paging"`
}

var postData *PostObject

func postJSONToStruct(response *http.Response){
     fmt.Println("Handling PostJSONToStruct")

     responseObject, err := ioutil.ReadAll(response.Body)
    handleError(err, "postContent.postJSONToStruct")

     postData = new (PostObject)

     err = json.Unmarshal([]byte(responseObject), &postData)
    handleError(err, "postContent.postJSONToStruct")

     //test struct
     fmt.Println("Test Post ID: " + postData.Data[0].ID)
}

