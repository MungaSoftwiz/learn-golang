package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"net/http"
)

//func Greet(writer *bytes.Buffer, name string) {

//If we use general purpose interface we can use it in
//both tests and in our application
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//We use the Greet() func inside cause the http.Res.Writer
//also implements io.Writer
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

/*
 * sets up an HTTP server listening on port 5001
 * registers the MyGreeterHandler function as the handler for all incoming HTTP requests.
 */
func main() {
	Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
