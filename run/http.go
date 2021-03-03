package run

import (
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redpatronus/httpret/service"
	"github.com/sqreen/go-agent/sdk/middleware/sqecho/v4"
)

func HttpServiceStart(svc *service.Svc) {

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: svc.Sentry.Dsn,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Then create your app
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Once it's done, you can attach the handler as one of your middleware
	app.Use(sentryecho.New(sentryecho.Options{}))

	if svc.Sqreen.Enabled == true {
		app.Use(sqecho.Middleware())
	}

	// Set up routes
	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "httpret (https://github.com/redpatronus/httpret)")
	})

	app.GET("/api/v1/browser", svc.GetBrowserDetails)
	app.GET("/api/v1/asn", svc.GetAsnDetails)
	app.GET("/api/v1/ipinfo", svc.GetIPInfo)
	app.GET("/api/v1/virustotal", svc.GetVirusTotalDetails)
	//app.GET("/api/v1/ssl", nil)

	// And run it
	app.Logger.Fatal(app.Start(svc.Listen))
}
