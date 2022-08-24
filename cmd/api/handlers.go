package main

import (
	"net/http"
)

type jsonResponse struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	// credentials são os dados que se espera receber da request
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJson(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid or missing json"
		_ = app.writeJson(w, http.StatusBadRequest, payload)
	}

	// TODO autenticar
	app.infoLog.Println(creds.Email, creds.Password)

	// response
	payload.Error = false
	payload.Message = "Success"

	err = app.writeJson(w, http.StatusOK, payload)
	if err != nil {
		app.infoLog.Println(err)
	}
}
