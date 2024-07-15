package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context();
	log.Println("Starting my request");
	defer log.Println("Request finished");

	select {
	case <-time.After(time.Second * 5):
		log.Println("Sucess on Request!");
		w.Write([]byte("Sucess!"));
	case <-ctx.Done():
		http.Error(w, "Request stopped", http.StatusRequestTimeout);
	}
}