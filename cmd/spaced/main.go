package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/brandur/spaced/endpoint"
	"github.com/brandur/spaced/middleware"
	"github.com/brandur/spaced/store"
	"github.com/brandur/spaced/store/memstore"
	"github.com/codegangsta/negroni"
	"github.com/heroku/rollrus"
	"github.com/joeshaw/envdecode"
	"github.com/phyber/negroni-gzip/gzip"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	conf, err := newConfFromEnv()
	if err != nil {
		panic(err)
	}

	initLog(conf.ForceTTY, conf.Source, conf.RollbarToken)

	st := initStore(conf.DatabaseURL)

	stack := buildMiddlewareStack(conf)
	stack.UseHandler(endpoint.BuildRouter(st))

	log.WithFields(log.Fields{
		"port": conf.Port,
	}).Infof("Starting web server on port %v", conf.Port)
	graceful.Run(":"+conf.Port, conf.GracefulRestartTimeout, stack)
}

//
// Types
//

// Conf provides the global app configuration mapped to environment variables
// through the use of envdecode. This should be handled only by main functions
// which will hand off more granular configuration structs to the various
// modules that they instantiate.
type Conf struct {
	// Postgres database URL. If not specified, we default to an in-memory
	// store.
	DatabaseURL string `env:"DATABASE_URL"`

	ForceTTY               bool          `env:"FORCE_TTY,default=false"`
	GracefulRestartTimeout time.Duration `env:"GRACEFUL_RESTART_TIMEOUT,default=10s"`
	Port                   string        `env:"PORT,required"`
	RollbarToken           string        `env:"ROLLBAR_TOKEN"`

	// Used for metric emission and as environment in Rollbar.
	Source string `env:"SOURCE,default=development"`
}

// Loads configuration from the process' environment.
func newConfFromEnv() (*Conf, error) {
	conf := &Conf{}
	err := envdecode.Decode(conf)
	return conf, err
}

//
// Helpers
//

func buildMiddlewareStack(conf *Conf) *negroni.Negroni {
	stack := negroni.New()

	// Recover from panics and send them up to Rollbar.
	stack.Use(middleware.NewRecovery())

	// Gzip-compress responses if requested.
	stack.Use(gzip.Gzip(gzip.DefaultCompression))

	return stack
}

// Initializes program logging. Configures errors to go to Rollbar if a token
// was specified.
func initLog(forceTTY bool, environment, rollbarToken string) {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: forceTTY,
	})

	// transmits to Rollbar on Error, Fatal, and Panic
	if environment != "" && rollbarToken != "" {
		rollrus.SetupLogging(rollbarToken, environment)

		log.WithFields(log.Fields{
			"environment": environment,
		}).Infof("Initialized Rollbar tracking with environment %v", environment)
	}
}

// Initializes a store for program information. If a database URL was given, a
// Postgres store is initialized, otherwise an in-memory store is used.
func initStore(databaseURL string) store.Store {
	if databaseURL == "" {
		st, err := memstore.NewMemstore()
		if err != nil {
			log.Fatal(err)
		}
		return st
	}

	log.Fatal("Postgres adapter not yet implemented")
	return nil
}
