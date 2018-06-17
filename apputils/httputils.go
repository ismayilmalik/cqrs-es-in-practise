package apputils

import (
	"encoding/json"
	"net/http"
)

func HTTPRespondWithError(w http.ResponseWriter, code int, message string) {
	HTTPRespondWithJSON(w, code, map[string]string{"error": message})
}

func HTTPRespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload != nil {
		var response []byte
		switch payload.(type) {
		case string:
			{
				str, _ := payload.(string)
				response = []byte(str)
			}
		default:
			response, _ = json.Marshal(payload)
		}
		w.Write(response)
	}
}
