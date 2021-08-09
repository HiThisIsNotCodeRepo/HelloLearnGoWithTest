package main

import (
	"log"
	"net/http"

	. "github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/9_Dependency_injection"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
