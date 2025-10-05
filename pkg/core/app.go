package core

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Application struct {
	Config   Config
	Db       *sql.DB
	Server   *http.ServeMux
	Router   RouterInterface
	Features []FeatureInterface
}

func NewDefault() *Application {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	a := &Application{
		Config: Config{
			App: AppConfig{
				Port: os.Getenv("APP_PORT"),
			},
			DB: DBConfig{
				Host:     os.Getenv("DB_HOST"),
				Username: os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASS"),
				Name:     os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				Driver:   os.Getenv("DB_DRIVER"),
			},
		},
	}

	return a
}

func (a *Application) Run() {
	log.Printf("Server running on port %s", a.Config.App.Port)

	defer a.Db.Close()

	a.LoadFeatures()
	a.LoadRoutes()

	err := http.ListenAndServe(fmt.Sprintf(":%s", a.Config.App.Port), a.Server)
	if err != nil {
		panic(err)
	}
}

func (a *Application) UseRouter(r RouterInterface) {
	a.Router = r
}

func (a *Application) LoadRoutes() {
	for _, r := range a.Router.GetRoutes() {
		newPath := fmt.Sprintf("%s %s", r.GetMethod(), r.GetPath())
		fmt.Println(newPath)
		handler := r.GetHandler()

		for _, m := range r.GetMiddlewares() {
			handler = m.RunMiddleware(handler)
		}

		a.Server.HandleFunc(newPath, r.GetHandler())
	}
}

func (a *Application) UseDatabase(db DatabaseInterface) {
	a.Db = db.GetInstance()
}

func (a *Application) UseHTTPServer(h HTTPServerInterface) {
	a.Server = h.GetInstance()
}

func (a *Application) AddFeature(f FeatureInterface) {
	a.Features = append(a.Features, f)
}

func (a *Application) LoadFeatures() {
	for _, f := range a.Features {
		f.Boot(a)
	}
}
