package eyeauth

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/class/pizza/logger"
	"github.com/class/pizza/storer"
	"github.com/justinas/nosurf"
	authboss "gopkg.in/authboss.v1"
	_ "gopkg.in/authboss.v1/auth"
	// _ "gopkg.in/authboss.v1/confirm"
	// _ "gopkg.in/authboss.v1/lock"

	_ "gopkg.in/authboss.v1/recover"
	_ "gopkg.in/authboss.v1/register"
	// _ "gopkg.in/authboss.v1/remember"
)

var snapshot *Eyeauth

type Eyeauth struct {
	AB *authboss.Authboss
}

// UserAuthenticator is something that can authorize that the current user is authentic
// and figure out who or what it is
// It can also update a user password
// And create its own router
// This is fully fulfilled by authboss.Authboss as of v0
type UserAuthenticator interface {
	CurrentUser(http.ResponseWriter, *http.Request) (interface{}, error)
	NewRouter() http.Handler
	Init(...string) error
	PasswordHasher
}

type PasswordHasher interface {
	UpdatePassword(w http.ResponseWriter, r *http.Request, ptPassword string, user interface{},
		updater func() error) error
}

type StorerSetter interface {
	SetStorer(s interface{})
}

func (e *Eyeauth) SetAuthStorer(s authboss.Storer) error {
	logger.Get().Println("Storer: ", s)
	e.AB.Storer = s
	// e.AB.RegisterStorer = s
	return e.AB.Init()
}

func Get() *Eyeauth {
	if snapshot == nil {
		New()
	}
	return snapshot
}

func isDisabled(ctx *authboss.Context) error {
	if ctx.User == nil {
		return authboss.ErrUserNotFound
	}

	if ctx.User["disabled"].(bool) {

		return errors.New("Test")
	}
	return nil
}
func New() error {
	e := &Eyeauth{}
	ab := authboss.New()
	storer.SetUpStores()
	ab.CookieStoreMaker = storer.NewCookieStorer
	ab.SessionStoreMaker = storer.NewSessionStorer
	ab.MountPath = "/"
	ab.AuthLoginOKPath = "/dash"
	ab.RootURL = "http://localhost:8080/dash"
	ab.ViewsPath = "templates"
	ab.XSRFName = "csrf_token"
	ab.Mailer = authboss.LogMailer(os.Stdout)
	ab.Policies = []authboss.Validator{
		authboss.Rules{
			FieldName:       "email",
			Required:        true,
			AllowWhitespace: false,
		},
		authboss.Rules{
			FieldName:       "password",
			Required:        true,
			MinLength:       4,
			MaxLength:       16,
			AllowWhitespace: false,
		},
	}
	ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}
	ab.Layout.Funcs(template.FuncMap{
		"Itoa": strconv.Itoa,
	})
	e.AB = ab
	snapshot = e
	return nil
}
