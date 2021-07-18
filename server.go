package main

import (
	"fmt"
)

// func main() {
// 	defaultPort := "8080"
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	router := mux.NewRouter()

// 	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello World")
// 	})

//     router.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

//     }).Methods(http.MethodPost)

// 	log.Println(http.ListenAndServe(":"+port, router))
// }

func main() {
	fmt.Println("Hello Go")
}
