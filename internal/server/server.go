package server

import (
	"log"
	"net/http"
	"simple-server/internal/handlers"
	"time"
)

func Start() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/people", handlers.PeopleHandler)
	http.HandleFunc("/health", handlers.HealthCheckHandler)

	go func() {
		err := http.ListenAndServe("localhost:8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.AfterFunc(10*time.Second, handlers.SendPostRequest)
	select {}
}
