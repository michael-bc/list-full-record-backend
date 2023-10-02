package http_server

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_http"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/configs"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/logger"
	"net/http"
	"time"
)

var (
	ServerAlreadyStarted = errors.New("server already started")
)

type HTTPServer struct {
	app       *echo.Echo
	address   string
	isStarted bool
}

func NewHTTPServer(
	cfg *configs.Config,
	log logger.Logger,
	version string,
	criterionController *brand_http.CriterionController,
) *HTTPServer {
	app := echo.New()

	app.Use(LoggerMiddleware(log))
	app.Use(middleware.Recover())
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.HTTPErrorHandler = ErrorHandler(log, cfg.Environment)

	app.GET("/", HealthHandler(cfg.ServiceName, version))

	criterionController.SetupRoutes(app)

	return &HTTPServer{
		app:     app,
		address: cfg.HTTPServerAddr,
	}
}

func (h *HTTPServer) Start(ctx context.Context) error {
	if h.isStarted {
		return ServerAlreadyStarted
	}

	h.isStarted = true
	defer func() {
		h.isStarted = false
	}()

	go func() {
		select {
		case <-ctx.Done():
			const shutdownTimeout = time.Second * 30

			ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()

			h.app.Shutdown(ctx)
		}
	}()

	if err := h.app.Start(h.address); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
