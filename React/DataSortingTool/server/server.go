package main

import(
    "fmt"
    "net/http"
)

func main(){
    fmt.Println("Setting Static Directory...")
    http.Handle("/", http.FileServer(http.Dir("static")))

    fmt.Println("Handling Server Functions...")
    http.HandleFunc("/FBToken", handleFBLogin)
    http.HandleFunc("/FBContent", handleFBContent)
    http.HandleFunc("/GraphData", handleGraphContent)

    fmt.Println("Starting Server...")
    e := http.ListenAndServe(":9000", nil)
    handleError(e, "server.ListenAndServer")
}
