package main

import(
    "net/http"
    "fmt"
    "io/ioutil"
    "strings"
    "io"
)

type ResponseObject struct{
    accessToken string
    longLivedToken string
    userID string
    expire string
}

var Data *ResponseObject

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
    decode = strings.Split(decode[1], "&expires=")

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
    dataRequest := respToString(response)
    if(err != nil){
        panic(err)
    }

    fmt.Println("User Posts: " + dataRequest)
    return dataRequest
}

func getUserVideos() string{
     fmt.Println("Handling User Videos")

     response, err  := http.Get("https://graph.facebook.com/" + Data.userID + "/videos/uploaded?access_token=" + Data.accessToken)
    dataRequest := respToString(response)
    if(err != nil){
        panic(err)
    }

    fmt.Println("User Videos: " + dataRequest)
    return dataRequest
}

func getUserStatus() string{
    fmt.Println("Handling User Status")

    //response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/feed?filter=app_2915120374/?access_token=" + Data.accessToken)
    response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/statuses?access_token=" + Data.accessToken)
    dataRequest := respToString(response)
    if(err != nil){
         panic(err)
    }

    fmt.Println("User Status: " + dataRequest)
    return dataRequest
}

func getUserPhotos() string{
     fmt.Println("Handling User Photos")

     response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/photos/uploaded?access_token=" + Data.accessToken)
     dataRequest := respToString(response)
     if(err != nil){
          panic(err)
     }

     fmt.Println("User Photos: " + dataRequest)
    return dataRequest
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
