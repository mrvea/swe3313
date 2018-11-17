package handler

import (
	"context"
	"errors"

	"github.com/class/pizza/env"
	"github.com/class/pizza/eyeauth"
	"github.com/class/pizza/model"
	"github.com/eyesore/httprouter"
	"github.com/justinas/nosurf"

	log "github.com/class/pizza/logger"

	"encoding/json"
	"net/http"
)

var (
	// ParsedJSON is the value used to retrieve request data from the request context
	ParsedJSON = jsonKey{}

	// LoggedInUser is the value used to retrieve the logged in user from the request context
	LoggedInUser = loggedInUserKey{}

	// ErrForbidden is thrown when authorization fails for a resource.
	ErrForbidden        = errors.New("You do not have permission to use this resource.")
	ErrMalformedRequest = errors.New("Your request was not formatted properly.")
)

type jsonKey struct{}
type loggedInUserKey struct{}

// Error represents and error that contains a status code.
type Error interface {
	error
	Status() int
}

// Router initially describes httprouter.Router, but is fulfilled by having a function "Handler"
type Router interface {
	Handler(string, string, http.Handler)
}

// HTTPError represents an error handling a route and contains an HTTP status code indicating the problem.
type HTTPError struct {
	Code int
	Err  error
}

func (he HTTPError) Error() string {
	return he.Err.Error()
}

// Status returns the status code that represents the HTTPError
func (he HTTPError) Status() int {
	return he.Code
}

type ForbiddenError struct {
	HTTPError
}

func NewForbiddenError() ForbiddenError {
	return ForbiddenError{
		HTTPError{http.StatusForbidden, ErrForbidden},
	}
}

type DefaultResponse struct {
	Success bool
}

// HandleFunc returns an error if something goes wrong.
type HandleFunc func(e *env.Env, w http.ResponseWriter, r *http.Request) error

// Handler is a wrapper for a handler function and an application context.
type Handler struct {
	*env.Env
	Pattern string
	Handle  HandleFunc
}

// New creates and returns a pointer to a new Handler.
func New(e *env.Env, p string, h HandleFunc) *Handler {
	ab := eyeauth.Get().AB
	// ab.LayoutDataMaker = layoutData
	ab.RootURL = e.BaseUrl
	return &Handler{e, p, h}
}

// Execute is an abstraction of ServeHTTP that allows errors that occur during typical
// requests to be handled in the same way as HandleFunc Errors
func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) error {
	var ctx context.Context
	if h.Env.AB == nil {
		log.Debug("AB is nil")
		ab := eyeauth.Get().AB
		h.Env.SetAuth(ab)
	}
	u, err := h.Env.AB.CurrentUser(w, r)
	if err != nil {
		return HTTPError{http.StatusInternalServerError, err}
	} else if u == nil {
		return NewForbiddenError()
	}
	ctx = r.Context()
	ctx = context.WithValue(ctx, LoggedInUser, u)
	r = r.WithContext(ctx)
	if r.Header.Get("Content-Type") == "application/json" {
		data, err := parseJSON(r)
		if err != nil {
			return HTTPError{http.StatusBadRequest, ErrMalformedRequest}
		}
		ctx = r.Context()
		ctx = context.WithValue(ctx, ParsedJSON, data)
		r = r.WithContext(ctx)
	}

	w.Header().Set("Cache-Control", "no-cache")
	return h.Handle(h.Env, w, r)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Execute(w, r)
	if r.URL.Path == "/dash" {
		return
	}
	if err != nil {
		switch e := err.(type) {
		case ForbiddenError:
			switch r.URL.Path {
			default:
				http.Redirect(w, r, "/login", http.StatusFound)
				// status := e.Status()
				// log.Infof("[ROUTE = %s %s] HTTP %d - %s", r.Method, h, status, e)
				// http.Error(w, e.Error(), status)
			}
		case HTTPError:
			status := e.Status()
			log.Infof("HTTP %d - %s", status, e)
			http.Error(w, e.Error(), status)
			// TODO tj - catch DB connection errors and ping a few times to see if it comes back. and retry?
		default:
			// 500
			log.Infof("[ROUTE = %s %s] Application Error : %s", r.Method, h, e)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (h *Handler) String() string {
	return h.Pattern
}

// Route creates a route on the router that uses the handler.
func (h *Handler) Route(method string, r Router) {
	r.Handler(method, h.Pattern, h)
}

func GetLoggedInUser(r *http.Request) *model.User {
	return r.Context().Value(LoggedInUser).(*model.User)
}

func GetRequestData(r *http.Request) map[string]interface{} {
	return r.Context().Value(ParsedJSON).(map[string]interface{})
}

func GetRouteParams(r *http.Request) httprouter.Params {
	return httprouter.ParamsFromContext(r.Context())
}

func parseJSON(r *http.Request) (map[string]interface{}, error) {
	var postData interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postData)
	if err != nil {
		log.Debug("Error parsing post data.")
		return nil, err
	}

	return postData.(map[string]interface{}), nil
}

func Nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(h) //this does nothing right now cause the success and failure handler are the same,
	return surfing
}
