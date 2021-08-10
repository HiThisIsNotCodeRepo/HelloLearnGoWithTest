package main

import (
	"log"
	"net/http"
	"os"

	. "github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/9_Dependency_injection"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	Greet(os.Stdout, "world")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
