package http

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/lokesh-go/go-boilerplate/api/rest/router"
	"github.com/lokesh-go/go-boilerplate/internal/dependencies"
	"github.com/lokesh-go/go-boilerplate/pkg/utils"
)

type httpServer struct {
	dep *dependencies.Deps
}

// New ...
func New(d *dependencies.Deps) *httpServer {
	return &httpServer{
		dep: d,
	}
}

// Start
func (s *httpServer) Start() (err error) {
	// Register routes
	router := router.New(s.dep)
	publicRouter := router.RegisterPublicRoutes()
	internalRouter := router.RegisterInternalRoutes()

	// Public server
	publicServer := &http.Server{
		Addr:         s.dep.Config.Get().Server.HTTP.PublicAddr,
		Handler:      publicRouter,
		ReadTimeout:  time.Duration(s.dep.Config.Get().Server.HTTP.ReadTimeoutInSeconds) * time.Second,
		WriteTimeout: time.Duration(s.dep.Config.Get().Server.HTTP.WriteTimeoutInSeconds) * time.Second,
	}

	// Internal server
	internalServer := &http.Server{
		Addr:         s.dep.Config.Get().Server.HTTP.InternalAddr,
		Handler:      internalRouter,
		ReadTimeout:  time.Duration(s.dep.Config.Get().Server.HTTP.ReadTimeoutInSeconds) * time.Second,
		WriteTimeout: time.Duration(s.dep.Config.Get().Server.HTTP.WriteTimeoutInSeconds) * time.Second,
	}

	// Global ctx that gets cancelled when the OS sends an interrupt/terminate signal (Ctrl+C)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create an error group with global context
	// If any goroutine in 'g' return an error, the entire gctx is cancelled
	// If ctx gets cancelled (SIGINT), gctx is cancelled too.
	g, gctx := errgroup.WithContext(ctx)

	// Start public server
	g.Go(func() error {
		// Log
		s.dep.Logger.Log.Info(fmt.Sprintf("public http server listening on %s", s.dep.Config.Get().Server.HTTP.PublicAddr))

		// Start the server
		if err := publicServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	// Start private server
	g.Go(func() error {
		// Log
		s.dep.Logger.Log.Info(fmt.Sprintf("internal http server listening on %s", s.dep.Config.Get().Server.HTTP.InternalAddr))

		// Start the server
		if err := internalServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	// Wait for the servers to finish
	g.Go(func() error {
		// Wait for the Signal or error
		// Any error or signal will cancel the gctx
		<-gctx.Done()

		// Log
		s.dep.Logger.Log.Info("shutting down servers ...")

		// Timeout for graceful shutdown
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dep.Config.Get().Server.HTTP.ShutdownTimeoutInSeconds)*time.Second)
		defer cancel()

		// Shutdown both servers
		var shutdownErr error
		if err := publicServer.Shutdown(shutdownCtx); err != nil {
			s.dep.Logger.Log.Error("failed to shutdown public server", utils.ErrorJSONKey, err)
			shutdownErr = err
		}
		if err := internalServer.Shutdown(shutdownCtx); err != nil {
			s.dep.Logger.Log.Error("failed to shutdown internal server", utils.ErrorJSONKey, err)
			shutdownErr = err
		}

		// Returns
		return shutdownErr
	})

	// Wait for the error group to finish
	if err := g.Wait(); err != nil {
		return err
	}

	// Returns
	return nil
}
