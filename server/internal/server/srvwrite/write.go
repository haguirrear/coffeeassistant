package srvwrite

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, body any) error {
	r, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(r); err != nil {
		return fmt.Errorf("error writing response: %w", err)
	}

	return nil
}
