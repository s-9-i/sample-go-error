package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	"sample-go-error/pkg/handler"
	"sample-go-error/pkg/infra"
	"sample-go-error/pkg/repository"
	"sample-go-error/pkg/service"
	"sample-go-error/pkg/usecase"
)

func main() {
	// Sentryクライアントの初期化
	initSentry()
	defer sentry.Flush(2 * time.Second)

	h := newHandler()
	e := echo.New()
	e.Use(errorMiddleware)

	e.GET("/return-internal-stack-error", h.ReturnInternalStackError)
	e.GET("/return-external-stack-error", h.ReturnExternalStackError)

	e.GET("/return-internal-callers-error", h.ReturnInternalCallersError)
	e.GET("/return-external-callers-error", h.ReturnExternalCallersError)

	e.Logger.Fatal(e.Start(":1323"))
}

func newHandler() *handler.Handler {
	d := infra.New()
	r := repository.New(d)
	s := service.New(r)
	u := usecase.New(s)
	return handler.New(u)
}

func errorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			fmt.Printf("%+v", err)
			c.Error(err)

			// Sentryにエラーを送信する処理
			ctx := c.Request().Context()
			hub := sentry.GetHubFromContext(ctx)
			if hub == nil {
				hub = sentry.CurrentHub().Clone()
				ctx = sentry.SetHubOnContext(ctx, hub)
			}
			hub.Scope().SetRequest(c.Request())
			hub.CaptureException(err)
		}
		return nil
	}
}

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_PUBLIC_DSN"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
