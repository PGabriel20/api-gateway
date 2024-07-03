package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    users := []User{
        {ID: 1, Name: "Foo"},
        {ID: 2, Name: "Bar"},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func main() {
    http.HandleFunc("/users", getUsers)
    log.Println("User Service running on port 8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
