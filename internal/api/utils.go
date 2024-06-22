package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func strToInt(text string) (int, error) {
	return strconv.Atoi(text)
}
