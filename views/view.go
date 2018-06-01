package views

import (
	"encoding/json"
	"net/http"
	"work/util"
)

// Render write bytes into http.ResponseWriter
func Render(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	util.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
