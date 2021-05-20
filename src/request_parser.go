package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type malformedRequest struct {
	status int
	msg string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func parseRequest(r *http.Request, dst interface{}) error {
	if r.Method != "POST" {
		return &malformedRequest{
			status: http.StatusMethodNotAllowed,
			msg: "Request method must be POST",
		}
	}

	if ! strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		return &malformedRequest{
			status: http.StatusUnsupportedMediaType,
			msg: "Content-Type header must be application/json",
		}
	}

	err := json.NewDecoder(r.Body).Decode(&dst)
	if err != nil {
		return &malformedRequest{
			status: http.StatusBadRequest,
			msg: err.Error(),
		}
	}

	return nil
}
