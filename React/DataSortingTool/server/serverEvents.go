package main

import(
    "fmt"
    "net/http"
    "io"
)

type ResponseObject struct{
    accessToken string
    longLivedToken string
    userID string
    expire string
}

var Data *ResponseObject

func handleFBLogin(w http.ResponseWriter, r *http.Request){
    fmt.Println("Handling Facebook Login Request...")

    Data = new (ResponseObject)

    response := objToString(r)

    Data.accessToken = getResponseItem(response, "accessToken")
    Data.userID = getResponseItem(response, "userID")
    Data.expire = getResponseItem(response, "expiresIn")
    Data.longLivedToken = getLongLivedUserAccessToken(Data.accessToken)

    fmt.Println("Access Token: " + Data.accessToken)
    fmt.Println("UserID: " + Data.userID)
    fmt.Println("Expire: " + Data.expire)
    fmt.Println("Long Lived Access Token: " + Data.longLivedToken)
}

func handleFBContent(w http.ResponseWriter, r *http.Request){
    fmt.Println("Handling Facebook Content Request...")

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
    //response, err := http.Post("localhost:9000/UI.html", "")
}

func handleGraphSort(w http.ResponseWriter, r *http.Request){
    fmt.Println("Handling Graph Data Request")

    r.ParseForm()
    dataRequest := r.FormValue("type")
    fmt.Println("Graph Type Request: " + dataRequest)

}
