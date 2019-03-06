package snipeapi

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func Handler(w http.ResponseWriter, r *http.Request) {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
                panic(err)
        }
        fmt.Fprintf(w, "Hi there, I love %s!\n", body)
}

