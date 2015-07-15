package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
)

func handleError(e error, errorType string){
    if(e != nil){
        fmt.Println("Error: " + errorType)
        panic(e)
    }
}

func getResponseItem(decode string, item string) string{
    responseItem := strings.SplitAfter(decode, ("%5B" + item + "%5D="))
    responseItem = strings.Split(responseItem[1], "&object")
    return responseItem[0]
}

func objToString(r *http.Request) string{
    body, err := ioutil.ReadAll(r.Body)
    handleError(err, "util.objToString")
    return string(body)
}

func respToString(resp *http.Response) string{
    body, err := ioutil.ReadAll(resp.Body)
    handleError(err, "util.respToString")
    return string(body)
}

