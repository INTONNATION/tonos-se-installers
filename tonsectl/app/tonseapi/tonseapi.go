package tonseapi

import (
    "fmt"
    "log"
    "net/http"
    //"github.com/gorilla/mux"
)

func tonseapi() {
    //myRouter := mux.NewRouter().StrictSlash(true)
    http.HandleFunc("/tonse/init", tonseInit)
    http.HandleFunc("/tonse/start", tonseStart)
    http.HandleFunc("/tonse/stop", tonseStop)
    http.HandleFunc("/tonse/status", tonseStatus)
    http.HandleFunc("/tonse/reset", tonseReset)
    http.HandleFunc("/tonse/upgrade", tonseUpgrade)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func tonseInit(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseInit")
}

func tonseStart(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseStart")
}

func tonseStop(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseStop")
}

func tonseStatus(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseStatus")
}

func tonseReset(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseReset")
}

func tonseUpgrade(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseUpgrade")
}

func RunApi() {
    tonseapi()
}
