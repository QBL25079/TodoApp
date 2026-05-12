package core_http_response

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

var StatusCodeUniitialized = -1

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode: StatusCodeUniitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(StatusCode int) {
	rw.ResponseWriter.WriteHeader(StatusCode)
	rw.statusCode = StatusCode
}

func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statusCode == StatusCodeUniitialized {
		panic("no status code set")
	}
	return rw.statusCode
}