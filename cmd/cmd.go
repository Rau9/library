package cmd

import (
	"fmt"
	"log"

	"github.com/Rau9/library/internal/books"
	"github.com/Rau9/library/internal/config"
	"github.com/Rau9/library/internal/healthcheck"
	lgr "github.com/Rau9/library/internal/logger"
	mw "github.com/Rau9/library/internal/middleware"
	"github.com/Rau9/library/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbclient  *gorm.DB
	zapLogger *zap.Logger
)

func init() {
	// initialise config as a singletone
	if err := config.Init(); err != nil {
		log.Fatalf("unable to intialize config: %s", err.Error())
	}
	// initialise server logger
	zl, err := lgr.NewProduction()
	if err != nil {
		log.Fatalf("unable to initialise logger: %s\n", err.Error())
	}
	zapLogger = zl
	// initialise dbclient
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetString("POSTGRES_HOST"),
		config.GetString("POSTGRES_USER"),
		config.GetString("POSTGRES_PASSWORD"),
		config.GetString("POSTGRES_DBNAME"),
		config.GetString("POSTGRES_PORT"),
		config.GetString("POSTGRES_SSLMODE"),
	)
	c, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dbclient = c
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}
}

func Start() {
	defer zapLogger.Sync()

	if err := storage.RunMigrations(dbclient); err != nil {
		panic("failed to run migrations")
	}

	e := echo.New()
	e.HideBanner = true

	// middlewares
	e.Use(mw.ZapLogger(zapLogger))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisableStackAll:   false,
		DisablePrintStack: false,
	}))

	e.File("/", "assets/index.html")
	e.File("/consult.html", "assets/consult.html")
	e.File("/create.html", "assets/create.html")
	e.File("/modify.html", "assets/modify.html")
	e.File("/remove.html", "assets/remove.html")
	e.Static("/css", "assets/css")
	e.Static("/jq", "assets/jq")
	e.Static("/img", "assets/img")
	e.Static("/js", "assets/js")
	e.GET("/healthz", healthcheck.Handler(dbclient, zapLogger))
	e.GET("/books", books.List(dbclient, zapLogger))
	e.GET("/books/:uuid", books.Read(dbclient, zapLogger))
	e.DELETE("/books/:uuid", books.Delete(dbclient, zapLogger))
	e.POST("/books", books.Create(dbclient, zapLogger))
	e.PUT("/books/:uuid", books.Update(dbclient, zapLogger))

	e.Logger.Fatal(e.Start(config.GetString("BIND_ADDRESS")))
}
