package handlers

import (
	"fmt"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "http web-server works correctly!")
}
