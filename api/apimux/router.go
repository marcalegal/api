package apimux

import (
	"strings"

	"github.com/gorilla/mux"
)

// Router provides just a convenience struct to store a base router
// and methods to specify differents API versions services.
type Router struct {
	multiplexer    *mux.Router
	versionRouters map[int]*Subrouter
}

// AddAPIVersion adds a new Subrouter instance constrained to the version
// specified if no version router exists, otherwise return the already created
// instance.
func (r *Router) AddAPIVersion(version int) *Subrouter {
	return r.apiVersionRouter(version)
}

func (r *Router) apiVersionRouter(version int) *Subrouter {
	if subrouter, ok := r.versionRouters[version]; ok {
		return subrouter
	}

	subrouter := r.multiplexer.
		Headers("Accept", AcceptHeader(version)).
		Subrouter()

	r.versionRouters[version] = &Subrouter{
		versionRouter:   subrouter,
		servicesRouters: make(map[string]*mux.Router),
	}

	return r.versionRouters[version]
}

// Multiplexer returns a pointer to the undelying mux.Router instance.
func (r *Router) Multiplexer() *mux.Router {
	return r.multiplexer
}

// AddService adds a service to a specific API version with a given prifix path.
func (r *Router) AddService(apiVersion int, prefixPath string, service Service) *Subrouter {
	subrouter := r.apiVersionRouter(apiVersion)
	subrouter.AddService(prefixPath, service)

	return subrouter
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	multiplexer := mux.NewRouter().
		PathPrefix("/api/").
		Subrouter()

	return &Router{
		multiplexer:    multiplexer,
		versionRouters: make(map[int]*Subrouter),
	}
}

// Subrouter stores all services of a specific API version.
type Subrouter struct {
	versionRouter   *mux.Router
	servicesRouters map[string]*mux.Router
}

func (s *Subrouter) serviceRouter(prefixPath string) *mux.Router {
	name := strings.Replace(prefixPath, "/", "", -1)

	if multiplexer, ok := s.servicesRouters[name]; ok {
		return multiplexer
	}

	multiplexer := s.versionRouter.
		StrictSlash(true).
		PathPrefix(prefixPath).
		Subrouter()
	s.servicesRouters[name] = multiplexer

	return multiplexer
}

// AddService add services to the underlying service mux.Router instance.
func (s *Subrouter) AddService(prefixPath string, service Service) *Subrouter {
	for _, route := range service {
		r := s.serviceRouter(prefixPath).
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)

		if q := route.Query; q != nil {
			r.Queries(q.Key, q.Value)
		}
	}

	return s
}
