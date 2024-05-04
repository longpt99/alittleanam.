// Package payload provides utilities for dealing with HTTP request and response payloads.
// It integrates with sibling packages log and errors.
package utils

import (
	"ala-coffee-search/utils/errs"
	"encoding/json"
	"log"
	"net/http"
)

type success struct {
	IsSuccess bool `json:"isSuccess"`
}

var (
	encodedErrResp []byte = json.RawMessage(`{"message":"Oops! Something went wrong. Please try again later or contact support for assistance."}`)
	IsSuccess             = success{true}
)

type response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Code   int         `json:"code"`
}

// ClientReporter provides information about an error such that client and
// server errors can be distinguished and handled appropriately.
type ClientReporter interface {
	error
	Message() map[string]string
	Status() int
}

// WriteError writes an appropriate error response to the given response
// writer. If the given error implements ClientReport, then the values from
// ErrorReport() and StatusCode() are written to the response, except in
// the case of a 5XX error, where the error is logged and a default message is
// written to the response.
func WriteError(w http.ResponseWriter, r *http.Request, e error) {
	if cr, ok := e.(ClientReporter); ok {
		status := cr.Status()
		if status >= http.StatusInternalServerError {
			handleInternalServerError(w, r, e)
			return
		}

		Write(w, r, cr.Message(), status)

		return
	}

	handleInternalServerError(w, r, e)
}

// Write writes the given payload to the response. If the payload
// cannot be marshaled, a 500 error is written instead. If the writer
// cannot be written to, then this function panics.
func Write(w http.ResponseWriter, r *http.Request, extras ...interface{}) {
	res := &response{
		Status: "success",
		Code:   http.StatusOK,
	}

	var payload interface{}

	for _, ex := range extras {
		switch t := ex.(type) {
		case int:
			res.Code = t
		default:
			payload = t
		}
	}

	if res.Code >= http.StatusBadRequest {
		res.Status = "error"
		if len(payload.(map[string]string)) == 0 {
			res.Error = http.StatusText(res.Code)
		} else {
			res.Errors = payload
		}
	} else {
		res.Data = payload
	}

	encoded, err := json.Marshal(res)
	if err != nil {
		handleInternalServerError(w, r, errs.E(err))
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if res.Code != http.StatusOK {
		w.WriteHeader(res.Code)
	}

	// w.Header().Add("Content-Encoding", "gzip")
	// gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	// if err != nil {
	// 	panic(errors.E(err))
	// }

	// defer gz.Close()

	// if _, err = gz.Write(encoded); err != nil {
	// 	panic(errors.E(err))
	// }

	_, err = w.Write(encoded)
	if err != nil {
		handleInternalServerError(w, r, errs.E(err))
		return
	}
}

// handleInternalServerError writes the given error to stderr and returns a
// 500 response with a default message.
func handleInternalServerError(w http.ResponseWriter, _ *http.Request, e error) {
	// log.AlarmWithContext(r.Context(), e)
	log.Printf("Err: %v", e)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if _, err := w.Write(encodedErrResp); err != nil {
		panic(errs.E(errs.Op("payload.handleInternalServerError"), err))
	}
}
