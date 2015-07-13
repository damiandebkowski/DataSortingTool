/*
    secret: 4fdb6bccf5ba741a71c898ab471e145c
    APP ID: 490045904486271
    redirect: http://localhost:9000/UI.html
*/

package main

import(
    "net/http"
    "fmt"
    "io/ioutil"
    "strings"
    //"os"
    //"bufio"
    //"strconv"
    //"log"
    //"encoding/json"
)

type ResponseObject struct{
    accessToken string
    userID string
    expire string
}

func getAccessToken(decode string) string{
    accessToken := strings.SplitAfter(decode, "%5BaccessToken%5D=")
    accessToken = strings.Split(accessToken[1], "&object")
    return accessToken[0]
}

func getUserID(decode string) string{
    userID := strings.SplitAfter(decode, "%5BuserID%5D=")
    userID = strings.Split(userID[1], "&object")
    return userID[0]
}

func getExpireTime(decode string) string{
    expire := strings.SplitAfter(decode, "%5BexpiresIn%5D=")
    expire = strings.Split(expire[1], "&object")
    return expire[0]
}

func storeData(data *ResponseObject){
    write := []byte("Access Token: " + data.accessToken + "\n" + "UserID: " + data.userID + "\n" + "Expire: " + data.expire + "\n")
    err := ioutil.WriteFile("data/data" + data.userID, write, 0644)
    if(err != nil){
        fmt.Println("Error: func storeData: Could Not Write To File")
        panic(err)
    }
}

func faceBookToken(w http.ResponseWriter, r *http.Request){
    fmt.Println("Getting Access Token")

    data := new (ResponseObject)

    body, err := ioutil.ReadAll(r.Body)

    if(err != nil){
        fmt.Println("Error: func faceBookToekn: Method Form Post Request: ioutil.ReadAll(r.Body)")
        panic(err)
    }

    response := string(body)
    //fmt.Println("String of Body: " + response)

    data.accessToken = getAccessToken(response)
    data.userID = getUserID(response)
    data.expire = getExpireTime(response)

    storeData(data)
}

func sortDataRequest(w http.ResponseWriter, r *http.Request){
    fmt.Println("Data Request Method: " + r.Method)
    if(r.Method == "POST"){
        r.ParseForm()
        gender := r.PostFormValue("Gender")
        age := r.FormValue("Age")
        location := r.FormValue("Location")
        education := r.FormValue("Education")
        fmt.Println("Gender: " + gender + " Age: " + age + " Location: " + location + " Education: " + education)
    }
}

func main(){
    fmt.Println("Starting Server...")
    http.Handle("/", http.FileServer(http.Dir("static")))
    http.HandleFunc("/SortData", sortDataRequest)
    http.HandleFunc("/FBToken", faceBookToken)
    http.ListenAndServe(":9000", nil)
}
