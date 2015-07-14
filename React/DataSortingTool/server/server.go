package main

import(
    "net/http"
    "fmt"
    "io/ioutil"
    "strings"
    "io"
    "encoding/json"
)

type ResponseObject struct{
    accessToken string
    longLivedToken string
    userID string
    expire string
}

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

func videoJSONToStruct(response *http.Response){
    fmt.Println("Handling VideoJSONToStruct")
    responseObject, err := ioutil.ReadAll(response.Body)
    if(err != nil){
        panic(err)
    }

    videoData = new (VideoObject)

    err = json.Unmarshal([]byte(responseObject), &videoData)
    if(err != nil){
         panic(err)
    }

    //test struct
    fmt.Println("Test Video ID: " + videoData.Data[0].ID)
}

func statusJSONToStruct(response *http.Response){
     fmt.Println("Handling StatusJSONToStruct")
     responseObject, err := ioutil.ReadAll(response.Body)
     if(err != nil){
         panic(err)
     }

     statusData = new (StatusObject)

     err = json.Unmarshal([]byte(responseObject), &statusData)
     if(err != nil){
        panic(err)
     }

     //test struct
    fmt.Println("Test Status ID: " + statusData.Data[0].ID)
}

func photoJSONToStruct(response *http.Response){
     fmt.Println("Handling PhotoJSONToStruct")
     responseObject, err := ioutil.ReadAll(response.Body)
     if(err != nil){
          panic(err)
     }

     photoData = new (PhotoObject)

     err = json.Unmarshal([]byte(responseObject), &photoData)
     if(err != nil){
         panic(err)
     }

     //test struct
     fmt.Println("Test Photo ID: " + photoData.Data[0].ID)
}

func postJSONToStruct(response *http.Response){
     fmt.Println("Handling PostJSONToStruct")
     responseObject, err := ioutil.ReadAll(response.Body)
     if(err != nil){
         panic(err)
     }

     postData = new (PostObject)

     err = json.Unmarshal([]byte(responseObject), &postData)
     if(err != nil){
          panic(err)
     }

     //test struct
     fmt.Println("Test Post ID: " + postData.Data[0].ID)
}

var Data *ResponseObject
var videoData *VideoObject
var statusData *StatusObject
var photoData *PhotoObject
var postData *PostObject

func getResponseItem(decode string, item string) string{
    responseItem := strings.SplitAfter(decode, ("%5B" + item + "%5D="))
    responseItem = strings.Split(responseItem[1], "&object")
    return responseItem[0]
}

func objToString(r *http.Request) string{
    body, err := ioutil.ReadAll(r.Body)

    if(err != nil){
        fmt.Println("Error: func objToString")
        panic(err)
    }

    return string(body)
}

func respToString(resp *http.Response) string{
    body, err := ioutil.ReadAll(resp.Body)

    if(err != nil){
        fmt.Println("Error: func repToString")
        panic(err)
    }

    return string(body)
}

func getLongLivedUserAccessToken(accessToken string) string{
    response, err := http.Get("https://graph.facebook.com/oauth/access_token?grant_type=fb_exchange_token&client_id=490045904486271&client_secret=4fdb6bccf5ba741a71c898ab471e145c&fb_exchange_token=" + accessToken)

    if(err != nil){
        fmt.Println("Error: func getLongLivedUserAccessToken")
        panic(err)
    }

    longLivedToken := respToString(response)
    decode := strings.SplitAfter(longLivedToken, "access_token=")

    return decode[0]
}

func handleFBLogin(w http.ResponseWriter, r *http.Request){
    fmt.Println("Handling Facebook Login Request")

    Data = new (ResponseObject)

    response := objToString(r)

    fmt.Println("Response Raw Data: " + response)

    Data.accessToken = getResponseItem(response, "accessToken")
    Data.userID = getResponseItem(response, "userID")
    Data.expire = getResponseItem(response, "expiresIn")
    Data.longLivedToken = getLongLivedUserAccessToken(Data.accessToken)

    fmt.Println("Access Token: " + Data.accessToken)
    fmt.Println("UserID: " + Data.userID)
    fmt.Println("Expire: " + Data.expire)
    fmt.Println("Long Lived Access Token: " + Data.longLivedToken)
}

func getUserPosts() string{
    fmt.Println("Handling User Posts")

    response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/posts/?access_token=" + Data.accessToken)
    if(err != nil){
        panic(err)
    }

    postJSONToStruct(response)

    return ""
}

func getUserVideos() string{
     fmt.Println("Handling User Videos")

     response, err  := http.Get("https://graph.facebook.com/" + Data.userID + "/videos/uploaded?access_token=" + Data.accessToken)
     if(err != nil){
         panic(err)
     }

     videoJSONToStruct(response)


    return ""
}

func getUserStatus() string{
    fmt.Println("Handling User Status")

    //response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/feed?filter=app_2915120374/?access_token=" + Data.accessToken)
    response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/statuses?access_token=" + Data.accessToken)
    if(err != nil){
         panic(err)
    }

    statusJSONToStruct(response)

    return ""
}

func getUserPhotos() string{
     fmt.Println("Handling User Photos")

     response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/photos/uploaded?access_token=" + Data.accessToken)
     if(err != nil){
          panic(err)
     }

     photoJSONToStruct(response)

    return ""
}

func handleFBContent(w http.ResponseWriter, r *http.Request){
    fmt.Println("Handling Facebook Content Request")

    r.ParseForm()
    dataRequest := r.FormValue("data")
    fmt.Println("Content Type Request: " + dataRequest)

    switch(dataRequest){
    case "posts":
        dataRequest = getUserPosts()
        break
    case "pictures":
        dataRequest = getUserPhotos()
        break
    case "videos":
        dataRequest = getUserVideos()
        break
    case "status":
        dataRequest = getUserStatus()
        break
    default:
        fmt.Println("Error HandleDataSort Switch Statement DataRequest")
    }

    io.WriteString(w, dataRequest)
}

func handleGraphSort(w http.ResponseWriter, r *http.Request){
    fmt.Println("Handling Graph Data Request")

    r.ParseForm()
    dataRequest := r.FormValue("type")
    fmt.Println("Graph Type Request: " + dataRequest)
}

func main(){
    fmt.Println("Starting Server...")
    http.Handle("/", http.FileServer(http.Dir("static")))
    http.HandleFunc("/FBToken", handleFBLogin)
    http.HandleFunc("/FBContent", handleFBContent)
    http.HandleFunc("/GraphData", handleGraphSort)
    http.ListenAndServe(":9000", nil)
}
