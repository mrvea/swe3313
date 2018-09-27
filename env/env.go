package env

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	_ "github.com/class/pizza/logger"
	"github.com/class/pizza/storer"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/nosurf"
	"github.com/kelseyhightower/envconfig"
	authboss "gopkg.in/authboss.v1"
	_ "gopkg.in/authboss.v1/auth"
	// _ "gopkg.in/authboss.v0/confirm"
	// _ "gopkg.in/authboss.v0/lock"
	_ "gopkg.in/authboss.v1/recover"
	// _ "gopkg.in/authboss.v0/register"
	// _ "gopkg.in/authboss.v0/remember"

	// load the driver
	mysql "github.com/go-sql-driver/mysql"

	"fmt"
)

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

// Env contains the daemon-scoped environment for the application.
type Env struct {
	ListenPort    string
	BaseUrl       string
	DataDirectory string
	// EnvironmentName is: production - development - staging - test
	EnvironmentName  string
	DB               *sqlx.DB
	AB               UserAuthenticator
	ConnectionString string
	Version          string
}

type env struct {
	Environment string
	BaseUrl     string `envconfig:"BASE_URL"`
	DataDir     string `envconfig:"DATA_DIR"`

	// webserver
	ListenPort string `default:"8080" envconfig:"LISTEN_PORT"`

	// database
	DbHost      string `default:"localhost" envconfig:"DB_HOST"`
	DbUsername  string `envconfig:"DB_USERNAME"`
	DbPassword  string `envconfig:"DB_PASSWORD"`
	DbPort      string `default:"3306" envconfig:"DB_PORT"`
	DbParseTime bool   `default:"true" envconfig:"DB_PARSE_TIME"`
	DbDatabase  string `envconfig:"DB_DATABASE" required:"true"`
}

// New creates and returns a pointer to a new Env, or an error if a critical configuration is missing.
func New(prefix string) (*Env, error) {
	e := &Env{}
	config := &env{}

	err := envconfig.Process(prefix, config)
	if err != nil {
		return nil, err
	}

	// db
	db := dbHelper{config.DbHost, config.DbUsername, config.DbPassword, config.DbPort,
		config.DbDatabase, config.DbParseTime}
	connection, err := db.connection()
	if err != nil {
		return nil, err
	}

	// version information comes from package.json of the frontend app, for now.
	// TODO actually read this from the environment
	versionPath := "./ui/package.json"
	packageJSON, err := ioutil.ReadFile(versionPath)
	packageDescription := make(map[string]interface{})
	if err == nil {
		err = json.Unmarshal(packageJSON, &packageDescription)
		if err == nil {
			e.Version = packageDescription["version"].(string)
		}
	}

	// create the env
	e.ListenPort = config.ListenPort
	e.BaseUrl = config.BaseUrl
	e.EnvironmentName = config.Environment
	e.DB = connection

	// e.AB = authBoss()
	if e.EnvironmentName != "production" {
		e.ConnectionString = db.connectionString()
	}

	return e, nil
}

func (e *Env) String() (s string) {
	s += fmt.Sprintln("") // newline
	s += fmt.Sprintln("Application Version: ", e.Version)
	s += fmt.Sprintln("ConnectionString: ", e.ConnectionString)
	s += fmt.Sprintln("ListenPort: ", e.ListenPort)
	s += fmt.Sprintln("EnvironmentName: ", e.EnvironmentName)
	s += fmt.Sprintln("Database connection: ", e.DB)
	return
}

func (e *Env) SetAuth(AB *authboss.Authboss) {
	e.AB = AB
}

func (e *Env) SetAuthStorer(s authboss.Storer) error {
	switch t := e.AB.(type) {
	case *authboss.Authboss:
		t.Storer = s
	default:
		// panic if it's not a storesetter
		ss := t.(StorerSetter)
		ss.SetStorer(s)
	}
	return e.AB.Init()
}

// database
type dbHelper struct {
	Host      string
	Username  string
	Password  string
	Port      string
	Database  string
	ParseTime bool
}

func (d dbHelper) connection() (*sqlx.DB, error) {
	connectionString := d.connectionString()
	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// still doesn't allow @ character in password
func (d dbHelper) connectionString() string {
	host := fmt.Sprintf("%s:%s", d.Host, d.Port)
	config := &mysql.Config{
		User:      d.Username,
		Passwd:    d.Password,
		Net:       "tcp", // TODO allow other
		Addr:      host,
		DBName:    d.Database,
		ParseTime: d.ParseTime,
	}

	return config.FormatDSN()
}

// quick hack for fixing connection string
func urlEncode(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func authBoss() *authboss.Authboss {
	ab := authboss.New()
	storer.SetUpStores()
	ab.CookieStoreMaker = storer.NewCookieStorer
	ab.SessionStoreMaker = storer.NewSessionStorer
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

	return ab
}
