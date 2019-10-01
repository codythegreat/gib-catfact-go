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
    resp, err := http.Get(apiAddress)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    catText := gjson.Get(string(body), "text")
    fmt.Println(catText.String())
}


