package main

import (
	"log"
	"magicball/handler"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fileServer))
	http.HandleFunc("/magicball", handler.Dispatch)
	http.HandleFunc("/health_check", handler.HealthCheck)
	http.HandleFunc("/monitor/l7check", handler.HealthCheck)

	log.Fatalln(http.ListenAndServe(":10780", nil))
}
