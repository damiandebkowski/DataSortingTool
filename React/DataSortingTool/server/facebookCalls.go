package main

import(
    "fmt"
    "net/http"
    "strings"
)

func getLongLivedUserAccessToken(accessToken string) string{
    fmt.Println("Requesting Long Lived Access Token...")
    response, err := http.Get("https://graph.facebook.com/oauth/access_token?grant_type=fb_exchange_token&client_id=490045904486271&client_secret=4fdb6bccf5ba741a71c898ab471e145c&fb_exchange_token=" + accessToken)

    handleError(err, "facebookCalls.getLongLivedUserAccessToken")

    longLivedToken := respToString(response)
    decode := strings.SplitAfter(longLivedToken, "access_token=")

    return decode[0]
}

func getUserPosts() string{
    fmt.Println("Handling User Posts...")

    response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/posts/?access_token=" + Data.accessToken)
    handleError(err, "facebookCalls.getUserPosts")

    postJSONToStruct(response)

    return "Test Load Posts"
}

func getUserVideos() string{
    fmt.Println("Handling User Videos...")

    response, err  := http.Get("https://graph.facebook.com/" + Data.userID + "/videos/uploaded?access_token=" + Data.accessToken)
    handleError(err, "facebookCalls.getUserVideos")
    videoJSONToStruct(response)

    /*
    data := ""

    for i := 0; i < len(videoData.Data); i++ {
        fmt.Println("Video ID: " + videoData.Data[i].ID)
        data += videoData.Data[i].ID + "/"
    }

    fmt.Println("Test Data: " + data)
    */
    return "Test Load Videos"
}

func getUserStatus() string{
    fmt.Println("Handling User Status...")

    response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/statuses?access_token=" + Data.accessToken)
    handleError(err, "facebookCalls.getUserStatus")

    statusJSONToStruct(response)
    /*
    data := ""
    indexSize := len(statusData.Data)
    for i:= 0; i < indexSize; i++ {
        fmt.Println("Status ID: " + statusData.Data[i].ID)
        data += videoData.Data[i].ID + "/"
    }
    */
    return "Test Load Status"
}

func getUserPhotos() string{
    fmt.Println("Handling User Photos...")

    response, err := http.Get("https://graph.facebook.com/" + Data.userID + "/photos/uploaded?access_token=" + Data.accessToken)
    handleError(err, "facebookCalls.getUserPhotos")

    photoJSONToStruct(response)

    return "Test Load Photos"
}

