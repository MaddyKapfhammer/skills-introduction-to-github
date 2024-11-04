package main

import "fmt";
import "net/http";
import "log";


func showGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nice to e-meet you!",);
}

func main() {
	http.HandleFunc("/api", showGreeting);
	log.Fatal(http.ListenAndServe(":8080", nil));
}