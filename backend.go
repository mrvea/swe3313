package main

import (
	"github.com/class/pizza/env"
	"github.com/class/pizza/handler"
	"github.com/class/pizza/logger"
	log "github.com/class/pizza/logger"
	"github.com/class/pizza/model"

	"flag"
	"fmt"
	"net/http"
	"os"
)

func failServerStart(err error) {
	logger.Get().Println("*** FATAL ERROR preventing server startup ***: ", err)
}

func startWebserver() {
	appEnv, err := env.New("pizza")
	if err != nil {
		logger.Get().Println("Failed to start the server")
		failServerStart(err)
		return
	}
	listenPort := ":" + appEnv.ListenPort
	model.SetDatasource(appEnv.DB)
	router, err := handler.Routes(appEnv)
	if err != nil {
		logger.Get().Println("Failed to handler.Routers", err)
		failServerStart(err)
	}

	if !model.DiscoverDatabase(appEnv.DB, 60) {
		logger.Get().Println("Failed to Discover database")
		failServerStart(model.ErrDatabaseConnectionFailed)
	}
	logger.Get().Println("***************** Starting Webserver on port " + listenPort +
		" ********************")
	err = http.ListenAndServe(listenPort, router)
	if err != nil {
		// TODO - notify someone
		// TODO prevent server restart loop
		failServerStart(err)
	}
}

func helloWorld() {
	log.Info("Hello World!")
}

func testConfig() {
	// database
	appEnv, err := env.New("pizza")
	if err != nil {
		log.Info("Unable to create application environment.")
		log.Info(err)
	}

	log.Info("This is the system environment: ")
	log.Info(os.Environ())

	log.Info("And the application environment: ")
	log.Info(appEnv)
}

func testRoutes() {
	appEnv, err := env.New("pizza")
	if err != nil {
		log.Info("Unable to create application environment.")
		log.Info(err)
	}

	_, err = handler.Routes(appEnv)
	// more probably it will just panic
	if err != nil {
		log.Info("Problem building router:")
		log.Info(err)
		os.Exit(1)
	}
	log.Info("Router built successfully!")
	// TODO potentially print out all routes - that would be cool!
}

func usage() {
	fmt.Printf("Usage: %s COMMAND [OPTIONS] \n", os.Args[0])
	fmt.Printf(`Commands:
		start - Run the webserver.
		hello - Check if the program works?  Remove this crap.
		test - Print out the application configuration.
		migrate - Run a database migration.
	`)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	rollback := flag.Bool("rollback", false, "If present with migrate command, roll back instead of forward.")
	howMany := flag.Int("n", 1, "If present with rollback option, rollback this many migrations.")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		args[0] = "exit"
	}
	command := args[0]

	switch command {
	case "start":
		startWebserver()
	case "stop":
		// TODO tj
		flag.Usage()
	case "restart":
		// TODO tj
		flag.Usage()
	case "reload":
		// TODO tj
		flag.Usage()
	case "hello":
		helloWorld()
	case "test":
		testConfig()
	case "migrate":
		err := model.Migrate(*rollback, *howMany)
		if err != nil {
			log.Info("Database migration failed with", err)
			os.Exit(1)
		}
	case "dbconnect":
		e, _ := env.New("pizza")
		gotDB := model.DiscoverDatabase(e.DB, 60)
		if !gotDB {
			os.Exit(1)
		}
	case "routes":
		testRoutes()
	default:
		flag.Usage()
	}
}
