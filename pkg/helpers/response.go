package helpers

import (
	"encoding/json"
	"net/http"
)

// JSONResponse - makes simple json response
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	res := map[string]interface{}{
		"status": status,
		"data":   data,
	}
	json.NewEncoder(w).Encode(res)
	w.Header().Set("Content-Type", "application/json")
}
