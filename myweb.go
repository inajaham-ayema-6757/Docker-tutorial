package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Friends ..!!!\n")
	fmt.Fprintf(w, "And I am Ameya Makarand Mahajani\n")
}

func main() {
	fmt.Println("Server is running ....")
	fmt.Println("Press Ctrl+C to close the server. ")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)

}
