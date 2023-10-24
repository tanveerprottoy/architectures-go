package modulelayer

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/content"
	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/user"
	modulerouter "github.com/tanveerprottoy/architectures-go/internal/modulelayer/router"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/config"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/constant"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/data"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/httpext"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/router"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/validatorext"
)

// App struct
type App struct {
	Server             *http.Server
	idleConnsClosed    chan struct{}
	DBClient           *data.Client
	HTTPClientProvider *httpext.ClientProvider
	router             *router.Router
	Middlewares        []any
	UserModule         *user.Module
	ContentModule      *content.Module
	Validate           *validator.Validate
}

// NewApp creates App
func NewApp() *App {
	a := new(App)
	a.initComponents()
	a.initServer()
	a.configureGracefulShutdown()
	return a
}

// initDB initializes DB client
func (a *App) initDB() {
	a.DBClient = data.GetInstance()
}

func (a *App) initRouter() {
	a.router = router.NewRouter()
}

func (a *App) initHTTPClientProvider() {
	a.HTTPClientProvider = httpext.NewClientProvider(90*time.Second, nil, nil)
}

// initValidator initializes validator
func (a *App) initValidator() {
	a.Validate = validator.New()
	validatorext.RegisterTagNameFunc(a.Validate)
	_ = a.Validate.RegisterValidation("notempty", validatorext.NotEmpty)
}

// initModules initializes application modules
func (a *App) initModules() {
	a.UserModule = user.NewModule(a.DBClient.DB, a.Validate)
	a.ContentModule = content.NewModule(a.DBClient.DB, a.Validate)
}

// initModuleRouters initializes module routers and routes
func (a *App) initModuleRouters() {
	modulerouter.RegisterUserRoutes(a.router, []string{constant.V1}, a.UserModule)
	modulerouter.RegisterContentRoutes(a.router, []string{constant.V1}, a.ContentModule)
}

// initServer initializes the server
func (a *App) initServer() {
	a.Server = &http.Server{
		Addr:    ":" + config.GetEnvValue("APP_PORT"),
		Handler: a.router.Mux,
	}
}

// configureGracefulShutdown configures graceful shutdown
func (a *App) configureGracefulShutdown() {
	// code to support graceful shutdown
	a.idleConnsClosed = make(chan struct{})
	go func() {
		// this func listens for SIGINT and initiates
		// shutdown when SIGINT is received
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		// We received an interrupt signal, shut down.
		log.Printf("Received an interrupt signal")
		if err := a.Server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP Server shutdown error: %v", err)
		}
		close(a.idleConnsClosed)
	}()
}

// ShutdownServer shuts down the server
func (a *App) ShutdownServer(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := a.Server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Println("Server shutdown")
		// add code
	}
}

// initComponents initializes application components
func (a *App) initComponents() {
	a.initDB()
	a.initRouter()
	a.initValidator()
	a.initModules()
	a.initModuleRouters()
}

// Run runs the server
func (a *App) Run() {
	// if err == http.ErrServerClosed do nothing
	if err := a.Server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP Server ListenAndServe: %v", err)
	}
	<-a.idleConnsClosed
	log.Println("Server shutdown")
}

// RunTLS runs the server with TLS
func (a *App) RunTLS() {
	// if err == http.ErrServerClosed do nothing
	if err := a.Server.ListenAndServeTLS("cert.crt", "key.key"); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTPS server ListenAndServe: %v", err)
	}
	<-a.idleConnsClosed
}

// RunListenAndServe runs the server
func (a *App) RunListenAndServe() {
	err := http.ListenAndServe(":"+config.GetEnvValue("APP_PORT"), a.router.Mux)
	if err != nil {
		panic(err)
	}
}

// RunListenAndServeTLS runs the server with TLS
func (a *App) RunListenAndServeTLS() {
	err := http.ListenAndServeTLS(":443", "cert.crt", "key.key", a.router.Mux)
	if err != nil {
		panic(err)
	}
}
