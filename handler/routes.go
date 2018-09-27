package handler

import (
	"net/http"

	"github.com/class/pizza/env"
	"github.com/class/pizza/eyeauth"
	"github.com/class/pizza/logger"
	log "github.com/class/pizza/logger"
	"github.com/class/pizza/model"
	"github.com/julienschmidt/httprouter"
)

// Routes creates and configures a mux for the domain specific http server.
func Routes(e *env.Env) (*httprouter.Router, error) {
	r := httprouter.New()

	// Configure Authenticator
	if err := eyeauth.Get().SetAuthStorer(model.UserTable()); err != nil {
		logger.Get().Println("Setting Storer")
		return nil, err
	}
	ab := Nosurfing(eyeauth.Get().AB.NewRouter())

	// authboss provided routes
	r.Handler("GET", "/login", ab)
	r.Handler("GET", "/logout", ab)
	r.Handler("POST", "/login", ab)
	r.Handler("GET", "/recover", ab)
	r.Handler("POST", "/recover", ab)
	r.Handler("GET", "/recover/complete", ab)
	r.Handler("POST", "/recover/complete", ab)
	r.Handler("GET", "/home", ab)

	// main logged in template
	// New(e, "/", UserHome).Route("GET", r)
	// New(e, "/dash/*uiroute", UserHome).Route("GET", r)
	New(e, "/home/*uiroute", UserHome).Route("GET", r)
	// users

	if e.EnvironmentName == "production" {
		r.PanicHandler = throw500
	}

	return r, nil
}

func throw500(w http.ResponseWriter, r *http.Request, p interface{}) {
	log.Info("Application panic on request: ", p)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
