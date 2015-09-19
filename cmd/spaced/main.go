package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/brandur/spaced"
	"github.com/brandur/spaced/middleware"
	"github.com/codegangsta/negroni"
	"github.com/heroku/rollrus"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/joeshaw/envdecode"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	conf, err := newConfFromEnv()
	if err != nil {
		panic(err)
	}

	initLog(conf.ForceTTY, conf.Source, conf.RollbarToken)

	server, err := spaced.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	// note that the "/" is special and handles all paths
	mux := http.NewServeMux()
	mux.Handle("/", server)

	stack := buildMiddlewareStack(conf)
	stack.UseHandler(mux)

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
