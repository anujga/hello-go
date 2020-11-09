package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

<<<<<<< HEAD:main.go
	fmt.Fprintf(w, "Morning:  "+title+"\n")
=======
	fmt.Fprintf(w, "Good Evening:  "+title+"\n")
>>>>>>> fd39bbe8b9f85387b95684113c8815b66253f0f5:pkg/main.go
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
