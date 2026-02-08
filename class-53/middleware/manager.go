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

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next // http.HandlerFunc(handlers.GetProducts)

	for _, middleware := range middlewares {
		n = middleware(n) // arekta(GetProducts)
	}

	// m.globalMiddlewares = [logger, hudai]
	// hudai(logger(http.handlerFunc(GetProducts)))
	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n) //Logger(arekta(GetProducts))
	}

	return n
}