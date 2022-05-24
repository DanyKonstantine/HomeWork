package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("OK")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}
