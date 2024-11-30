package services

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/funcmap"
	"github.com/mikestefanello/pagoda/pkg/seeders"

	// Require by ent
	_ "github.com/mikestefanello/pagoda/ent/runtime"
)

// Container contains all services used by the application and provides an easy way to handle dependency
// injection including within tests
type Container struct {
	// Validator stores a validator
	Validator *Validator

	// Web stores the web framework
	Web *echo.Echo

	// Config stores the application configuration
	Config *config.Config

	// Cache contains the cache client
	Cache *CacheClient

	// Database stores the connection to the database
	Database *sql.DB

	// ORM stores a client to the ORM
	ORM *ent.Client

	// Mail stores an email sending client
	Mail *MailClient

	// Auth stores an authentication client
	Auth *AuthClient

	// TemplateRenderer stores a service to easily render and cache templates
	TemplateRenderer *TemplateRenderer

	// Tasks stores the task client
	// Tasks *backlite.Client
}

// NewContainer creates and initializes a new Container
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initValidator()
	c.initWeb()
	c.initCache()
	c.initDatabase()
	c.initORM()
	c.initAuth()
	c.initTemplateRenderer()
	c.initMail()
	// c.initTasks()
	return c
}

// Shutdown shuts the Container down and disconnects all connections.
// If the task runner was started, cancel the context to shut it down prior to calling this.
func (c *Container) Shutdown() error {
	if err := c.ORM.Close(); err != nil {
		return err
	}
	if err := c.Database.Close(); err != nil {
		return err
	}
	c.Cache.Close()

	return nil
}

// initConfig initializes configuration
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg

	// Configure logging
	switch cfg.App.Environment {
	case config.EnvProduction:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

// initValidator initializes the validator
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// initWeb initializes the web framework
func (c *Container) initWeb() {
	c.Web = echo.New()
	c.Web.HideBanner = true
	c.Web.Validator = c.Validator
}

// initCache initializes the cache
func (c *Container) initCache() {
	store, err := newInMemoryCache(c.Config.Cache.Capacity)
	if err != nil {
		panic(err)
	}

	c.Cache = NewCacheClient(store)
}

// initDatabase initializes the database
func (c *Container) initDatabase() {
	var err error
	var connection string

	switch c.Config.App.Environment {
	case config.EnvTest:
		connection = c.Config.Database.TestConnection
	default:
		connection = c.Config.Database.Connection
	}

	c.Database, err = openDB("postgres", connection)
	if err != nil {
		panic(fmt.Sprintf("failed to open database: %v", err))
	}
}

// initORM initializes the ORM
func (c *Container) initORM() {
	drv := entsql.OpenDB(c.Config.Database.Driver, c.Database)

	c.ORM = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	ctx := context.Background()
	if err := c.ORM.Schema.Create(ctx); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	// Run seeders
	if err := seeders.SeedAll(c.ORM); err != nil {
		panic(fmt.Sprintf("failed to run seeders: %v", err))
	}
}

// initAuth initializes the authentication client
func (c *Container) initAuth() {
	c.Auth = NewAuthClient(c.Config, c.ORM)
}

// initTemplateRenderer initializes the template renderer
func (c *Container) initTemplateRenderer() {
	c.TemplateRenderer = NewTemplateRenderer(c.Config, c.Cache, funcmap.NewFuncMap(c.Web))
}

// initMail initialize the mail client
func (c *Container) initMail() {
	var err error
	c.Mail, err = NewMailClient(c.Config, c.TemplateRenderer)
	if err != nil {
		panic(fmt.Sprintf("failed to create mail client: %v", err))
	}
}

// openDB opens a database connection
func openDB(driver, connection string) (*sql.DB, error) {
	// Helper to automatically create the directories that the specified sqlite file
	// should reside in, if one
	if driver == "sqlite3" {
		d := strings.Split(connection, "/")

		if len(d) > 1 {
			path := strings.Join(d[:len(d)-1], "/")

			if err := os.MkdirAll(path, 0755); err != nil {
				return nil, err
			}
		}
	}

	return sql.Open(driver, connection)
}
