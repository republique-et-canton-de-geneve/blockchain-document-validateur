package restapi

import (
	"context"
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	swag "github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	internal "github.com/Magicking/rc-ge-ch-pdf/internal"
	"github.com/Magicking/rc-ge-ch-pdf/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../docs/rc-ge-ch.yml

var ethopts struct {
	WsURI      string `long:"ws-uri" env:"WS_URI" description:"Ethereum WS URI (e.g: ws://HOST:8546)"`
	PrivateKey string `long:"pkey" env:"PRIVATE_KEY" description:"hex encoded private key"`
}

var serviceopts struct {
	DbDSN string `long:"db-dsn" env:"DB_DSN" description:"Database DSN (e.g: /tmp/test.sqlite)"`
}

func configureFlags(api *operations.RCGHorodatageAPI) {
	ethOpts := swag.CommandLineOptionsGroup{
		LongDescription:  "",
		ShortDescription: "Ethereum options",
		Options:          &ethopts,
	}
	serviceOpts := swag.CommandLineOptionsGroup{
		LongDescription:  "",
		ShortDescription: "Service options",
		Options:          &serviceopts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ethOpts, serviceOpts}
}

func configureAPI(api *operations.RCGHorodatageAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	ctx := internal.NewDBToContext(context.Background(), serviceopts.DbDSN)
	ctx = internal.NewCCToContext(ctx, ethopts.WsURI)
	ctx = internal.NewBLKToContext(ctx, ethopts.WsURI, ethopts.PrivateKey)

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BinProducer = runtime.ByteStreamProducer()

	api.GetreceiptHandler = operations.GetreceiptHandlerFunc(func(params operations.GetreceiptParams) middleware.Responder {
		return internal.GetreceiptHandler(ctx, params)
	})
	api.ListtimestampedHandler = operations.ListtimestampedHandlerFunc(func(params operations.ListtimestampedParams) middleware.Responder {
		return internal.ListtimestampedHandler(ctx, params)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(ctx, api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	return internal.ValidateHandler(ctx, "/validate", internal.UploadHandler(ctx, "/upload", handler))
}
