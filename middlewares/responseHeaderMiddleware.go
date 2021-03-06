package middlewares

import "net/http"

type ResponseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

//NewResponseHeader constructs a new ResponseHeader middleware handler
func NewResponseHeader(handlerToWrap http.Handler, headerName, headerValue string) *ResponseHeader {
	return &ResponseHeader{handler: handlerToWrap, headerName: headerName, headerValue: headerValue}
}

//ServeHTTP handles the request by adding the response header
func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(rh.headerName, rh.headerValue)
	rh.handler.ServeHTTP(w, r)
}
