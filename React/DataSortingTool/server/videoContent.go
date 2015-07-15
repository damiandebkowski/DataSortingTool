package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type VideoObject struct {
	Data []struct {
		Comments struct {
			Data []struct {
				CanRemove   bool   `json:"can_remove"`
				CreatedTime string `json:"created_time"`
				From        struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"from"`
				ID        string `json:"id"`
				LikeCount int    `json:"like_count"`
				Message   string `json:"message"`
				UserLikes bool   `json:"user_likes"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					After  string `json:"after"`
					Before string `json:"before"`
				} `json:"cursors"`
			} `json:"paging"`
		} `json:"comments"`
		CreatedTime string `json:"created_time"`
		Description string `json:"description"`
		EmbedHTML   string `json:"embed_html"`
		Format      []struct {
			EmbedHTML string `json:"embed_html"`
			Filter    string `json:"filter"`
			Height    int    `json:"height"`
			Picture   string `json:"picture"`
			Width     int    `json:"width"`
		} `json:"format"`
		From struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		Icon        string `json:"icon"`
		ID          string `json:"id"`
		Picture     string `json:"picture"`
		Published   bool   `json:"published"`
		Source      string `json:"source"`
		UpdatedTime string `json:"updated_time"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			After  string `json:"after"`
			Before string `json:"before"`
		} `json:"cursors"`
	} `json:"paging"`
}

var videoData *VideoObject

func videoJSONToStruct(response *http.Response){
    fmt.Println("Handling VideoJSONToStruct...")

    responseObject, err := ioutil.ReadAll(response.Body)
    handleError(err, "videoContent.videoJSONToStruct")

    videoData = new (VideoObject)

    err = json.Unmarshal([]byte(responseObject), &videoData)
    handleError(err, "videoContent.videoJSONToStruct")
}

