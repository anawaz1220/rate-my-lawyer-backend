package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/services"
)

// BuildRouter builds the router
func BuildRouter(c *services.Container) error {
	// Static files with proper cache control
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Create a group for all routes
	g := c.Web.Group("")

	g.Use(echomw.CORS())
	// Add CORS middleware
	g.Use(echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://yourdomain.com", "file:///D:/Personel/Freelance/Farooq/Tanvi/RateMyLawyer/Frontend/index.html"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	// Middleware
	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		middleware.SetLogger(),
		middleware.LogRequest(),
		echomw.Gzip(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		middleware.Session(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		middleware.LoadAuthenticatedUser(c.Auth),
		middleware.ServeCachedPage(c.TemplateRenderer),
	)

	// Initialize and register all handlers
	for _, h := range GetHandlers() {
		if err := h.Init(c); err != nil {
			return err
		}
		h.Routes(g)
	}

	// Error handler
	errHandler := Error{c.TemplateRenderer}
	c.Web.HTTPErrorHandler = errHandler.Page

	return nil
}
