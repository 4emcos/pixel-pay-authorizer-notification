package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var random = rand.New(src)

func main() {
	http.HandleFunc("/authorization", authorizationHandler)
	http.HandleFunc("/notification", notificationHandler)
	port := "8099"
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func authorizationHandler(w http.ResponseWriter, r *http.Request) {
	weightedRandom := random.Intn(100) < 80
	json.NewEncoder(w).Encode(map[string]bool{"authorization": weightedRandom})
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	weightedRandom := random.Intn(100) < 80
	json.NewEncoder(w).Encode(map[string]bool{"notification": weightedRandom})
}
