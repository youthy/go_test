/*Package views provide functions convert struct to show data */
package views

import (
	"encoding/json"
	"go_test/util"
	"net/http"
)

// Render write bytes into http.ResponseWriter
func Render(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	util.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
