package main

import (
	"encoding/json"
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

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// error response
		app.errorLog.Println("Invalid json")
		payload.Error = true
		payload.Message = "Invalid json"

		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(out)
		return
	}

	// TODO autenticar
	app.infoLog.Println(creds.Email, creds.Password)

	// response
	payload.Error = false
	payload.Message = "Success"

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.infoLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
