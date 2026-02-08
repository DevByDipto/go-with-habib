package middleware

import "net/http"

// Middleware type definition
type Middleware func(http.Handler) http.Handler

// Manager struct to hold global middlewares
type Manager struct {
    globalMiddlewares []Middleware
}

// NewManager initializes a new Manager
func NewManager() *Manager {
    return &Manager{
        globalMiddlewares: make([]Middleware, 0),
    }
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...) 
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler 

	for _, middleware := range middlewares {
		h = middleware(h) 
	}

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler 

	
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h) 
	}

	return h
}