package controller

import (
	"fmt"
	"net/http"
)

func StartServer() error {
	if err := http.ListenAndServe(SERVER_PORT, nil); err != nil {
		return fmt.Errorf("Server error: %s", err)
	}

	return nil
}
