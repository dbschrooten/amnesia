package helpers

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, v interface{}, opt ...map[string]interface{}) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	enc := json.NewEncoder(w)

	if len(opt) > 0 {
		if v, ok := opt[0]["StatusCode"].(int); ok {
			w.WriteHeader(v)
		}

		if v, ok := opt[0]["PrettyPrint"].(bool); ok {
			if v {
				enc.SetIndent("", "    ")
			}
		}
	}

	if err := enc.Encode(v); err != nil {
		return err
	}

	return nil
}
