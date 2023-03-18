package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) readJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// limiting the size of the json file
	myBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(myBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		_ = errors.New("request body must have only a single json file")
	}

	return nil
}

func (app *application) writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil
}

func (app *application) errorJson(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	app.writeJson(w, statusCode, payload)
}
