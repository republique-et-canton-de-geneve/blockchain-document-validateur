package restapi

import (
	"context"
	"crypto/tls"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	swag "github.com/go-openapi/swag"
	"net/http"
//	"strings"
	internal "github.com/geneva_validateur/internal"
	operations "github.com/geneva_validateur/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --spec ../docs/rc-ge-validator.yml

var ethopts struct {
	WsURI         string `long:"ws-uri" env:"WS_URI" description:"Ethereum WS URI (e.g: ws://HOST:8546)"`
	LockedAddress string `long:"locked-addr" env:"LOCKED_ADDR" description:"Ethereum address of the sole verifier (anchor emitter)"`
}

func configureFlags(api *operations.RCGHorodatageValidateurAPI) {
	ethOpts := swag.CommandLineOptionsGroup{
		LongDescription:  "",
		ShortDescription: "Ethereum options",
		Options:          &ethopts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ethOpts}
}

func configureAPI(api *operations.RCGHorodatageValidateurAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	ctx := internal.NewCCToContext(context.Background(), ethopts.WsURI)
	ctx = internal.NewBLKToContext(ctx, ethopts.WsURI)
	ctx = internal.NewMonitoringToContext(ctx, ethopts.WsURI, ethopts.LockedAddress)
	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetStatusHandler = operations.GetStatusHandlerFunc(func(params operations.GetStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetStatus has not yet been implemented")
	})
	api.MonitoringHandler = operations.MonitoringHandlerFunc(func(params operations.MonitoringParams) middleware.Responder {
		return internal.MonitoringHandler(ctx, params)
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
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	//addr := strings.TrimPrefix(ethopts.LockedAddress, "0x")
	//return internal.ValidateHandler(ctx, "/validate", addr, handler)
	return internal.ValidateHandler(ctx, "/validate", ethopts.LockedAddress, handler)
}
