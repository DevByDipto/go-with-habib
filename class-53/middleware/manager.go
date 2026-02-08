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
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...) // [Logger(next),Hudai(next)]
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler // http.HandlerFunc(handlers.GetProducts)

	for _, middleware := range middlewares {
		h = middleware(h) // arekta(GetProducts)
	}

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler 

	
	// m.globalMiddlewares = [logger, hudai]
	// hudai(logger(http.handlerFunc(GetProducts)))
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h) //Logger(arekta(GetProducts))
	}

	return h
}