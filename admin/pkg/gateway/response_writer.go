package gateway

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	status int
}
