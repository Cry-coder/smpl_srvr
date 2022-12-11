package http

import (
	"context"
	"fmt"
	"github.com/Cry-coder/smpl_srvr/internal/infra/http/controllers"
	"net/http"
	"time"
)

func Session(next http.Handler) http.Handler {
	return controllers.SessionManager.LoadAndSave(next)
}
func Server(
	ctx context.Context,
	router http.Handler,
) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8007),
		Handler: router,
	}

	errServeCh := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errServeCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("error during server shutdown: %w", err)
		}
	case err := <-errServeCh:
		return fmt.Errorf("error during server execution: %w", err)
	}
	return nil
}
