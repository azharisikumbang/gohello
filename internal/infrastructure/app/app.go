package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/azharisikumbang/gohello/internal/infrastructure/database"
	"github.com/azharisikumbang/gohello/internal/infrastructure/server"
	"github.com/joho/godotenv"
)

type Application struct {
	Config Config
	DB     *sql.DB
	Server *http.ServeMux
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

func (app *Application) Run() {
	log.Printf("Server running on port %s", app.Config.App.Port)

	app.LoadRoutes()

	err := http.ListenAndServe(fmt.Sprintf(":%s", app.Config.App.Port), app.Server)
	if err != nil {
		panic(err)
	}
}

func NewDefault() *Application {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &Application{
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
			},
		},
	}

	var mysql *database.MySQL = &database.MySQL{
		Host:     app.Config.DB.Host,
		Username: app.Config.DB.Username,
		Password: app.Config.DB.Password,
		Name:     app.Config.DB.Name,
		Port:     app.Config.DB.Port,
	}

	app.DB = database.CreateDatabaseInstance(mysql).GetInstance()
	app.Server = server.CreateHTTPServer()

	return app
}

func (app *Application) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	newPath := fmt.Sprintf("GET %s", path)
	app.Server.HandleFunc(newPath, handler)
}

func (app *Application) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	newPath := fmt.Sprintf("POST %s", path)
	app.Server.HandleFunc(newPath, handler)
}

func (app *Application) AddDBContext(db *sql.DB) {
	app.DB = db
}
