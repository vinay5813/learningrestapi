package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"strconv"
	"vinay/test/config"
	"vinay/test/handler"
	"vinay/test/platform"
)

var AppVersion string

// This is a main.
//
// @contact.name vinay.yadav
// @host
// @BasePath /dummy_project/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization.
func main() {
	// Get config amd load configuration
	conf, key, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("bad configuration")
	}

	db := sqlx.MustConnect(conf.DBType, conf.DBDatasource)

	fmt.Println(db)

	e := echo.New()
	e.Validator = platform.NewValidator()
	e.HideBanner = true

	const (
		v1Path = "/vinay_test/v1"

		createProviderRoute = "/provider"

		pingRoute = v1Path + "/service/monitoring/ping"
	)

	e.GET(pingRoute, handler.Ping(AppVersion)).Name = "ping"
	e.GET(createProviderRoute, handler.ProviderGet(db))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.AppPort)))

	fmt.Println(conf, key)

}
