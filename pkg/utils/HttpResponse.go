package utils

import "net/http"

type HttpResponse struct {
	w http.ResponseWriter
	r *http.Request
}

func (r *HttpResponse) Response() error {
	return nil
}
