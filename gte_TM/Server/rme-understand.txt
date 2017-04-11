To understand web dev
you must understand these core concepts:

INTERFACE
defining functionality; defining behavior

REQUEST RESPONSE patter

http.Handler
----------------------
type Handler interface {
    SeverHTTP(ResponseWriter, *Request)
}
----------------------
A HANDLER will handle a request by providing a response.

web programming synonymous terms:
router
request router
multiplexer
mux
servemux
server
