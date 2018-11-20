package respond

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	error_message := map[string]interface{}{
		"Message": message,
		"Status":  0,
	}
	RespondWithJSON(w, code, error_message)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if code != 400{
		payload = map[string]interface{}{
			"Message" : "Success",
			"Status"  : 1,
			"Data"    : payload,
		}
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
