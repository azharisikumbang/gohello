package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	core "github.com/azharisikumbang/gohello/internal"
	"github.com/joho/godotenv"
)

type Application struct {
	Config   core.Config
	Db       core.DatabaseInterface
	Server   core.HTTPServerInterface
	Router   core.RouterInterface
	Features []core.FeatureInterface
}

func (a *Application) GetHTTPServer() core.HTTPServerInterface {
	return a.Server
}

func (a *Application) GetRouter() core.RouterInterface {
	return a.Router
}

func (a *Application) GetDatabase() core.DatabaseInterface {
	return a.Db
}

func (a *Application) GetLogger() {
}

func (a *Application) GetConfig() core.Config {
	return a.Config
}

func NewDefault() core.ApplicationInterface {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Application{
		Config: core.Config{
			App: core.AppConfig{
				Port: os.Getenv("APP_PORT"),
				Key:  os.Getenv("APP_KEY"),
			},
			DB: core.DBConfig{
				Host:     os.Getenv("DB_HOST"),
				Username: os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASS"),
				Name:     os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				Driver:   os.Getenv("DB_DRIVER"),
			},
		},
	}
}

func (a *Application) Run() {
	log.Printf("Server running on port %s", a.Config.App.Port)

	defer a.Db.GetInstance().Close()

	a.LoadFeatures()
	a.LoadRoutes()

	err := http.ListenAndServe(fmt.Sprintf(":%s", a.Config.App.Port), a.Server.GetInstance())
	if err != nil {
		panic(err)
	}
}

func (a *Application) UseRouter(r core.RouterInterface) {
	a.Router = r
}

func (a *Application) LoadRoutes() {
	mux := a.Server.GetInstance()
	for _, r := range a.Router.GetRoutes() {
		newPath := fmt.Sprintf("%s %s", r.GetMethod(), r.GetPath())
		handler := r.GetHandler()

		for _, m := range r.GetMiddlewares() {
			handler = m.RunMiddleware(handler)
		}

		mux.HandleFunc(newPath, handler)
	}
}

func (a *Application) UseDatabase(db core.DatabaseInterface) {
	if db.GetInstance() == nil {
		panic("Error: Database intance return nil.")
	}

	a.Db = db
}

func (a *Application) UseHTTPServer(h core.HTTPServerInterface) {
	if h.GetInstance() == nil {
		panic("Error: HTTP Server instance return nil.")
	}

	a.Server = h
}

func (a *Application) AddFeature(f core.FeatureInterface) {
	a.Features = append(a.Features, f)
}

func (a *Application) LoadFeatures() {
	for _, f := range a.Features {
		f.Boot(a)
	}
}
