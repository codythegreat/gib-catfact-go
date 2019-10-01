// application that gives random cat facts when called from the command line

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "github.com/tidwall/gjson"
)

// api web address
var apiAddress string = "https://cat-fact.herokuapp.com/facts/random"


func main() {
    // make a get request to the api
    resp, err := http.Get(apiAddress)
    if err != nil {
        log.Fatal(err)
    }
    // store the body of the resp
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    // use gjson to select the "text" value from the json 
    catText := gjson.Get(string(body), "text")
    fmt.Println(catText.String())
}


