package main

import (
	"net/http"
)

type Status string

const (
	StatusAvailable   Status = "available"
	StatusUnavailable Status = "unavailable"
)

type Health struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type EnvelopHealth struct {
	Status     Status `json:"status"`
	SystemInfo Health `json:"system_info"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": StatusAvailable,
		"system_info": Health{
			Environment: app.config.env,
			Version:     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
