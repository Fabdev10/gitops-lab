package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
)

type Comment struct {
    ID     int    `json:"id"`
    PostID int    `json:"post_id"`
    Author string `json:"author"`
    Text   string `json:"text"`
}

var comments = []Comment{
    {ID: 1, PostID: 1, Author: "alice", Text: "First comment from the Go service."},
    {ID: 2, PostID: 1, Author: "bob", Text: "Another sample comment for learning."},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "api-go"})
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(comments)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/health", healthHandler)
    mux.HandleFunc("/comments", commentsHandler)

    log.Printf("listening on :%s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatal(err)
    }
}
