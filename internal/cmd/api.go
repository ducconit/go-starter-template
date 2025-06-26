package cmd

import (
	"app/db"
	"app/internal/api"
	"app/internal/api/middleware"
	appConfig "app/internal/config"
	"context"
	"core/config"
	"core/log"
	"core/util"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API server",
	Long:  `API server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		listen, _ := cmd.Flags().GetString("listen")
		var host, port string
		if listen != "" {
			if !strings.Contains(listen, ":") {
				host = ""
				port = listen
			} else {
				host = listen[:strings.Index(listen, ":")]
				port = listen[strings.Index(listen, ":")+1:]
			}
			config.Set("api.host", host)
			config.Set("api.port", port)
		}
		// Set mode
		// Set environment
		isProduction := os.Getenv("APP_ENV") == "production"
		if isProduction {
			gin.SetMode(gin.ReleaseMode)
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		} else {
			gin.SetMode(gin.DebugMode)
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		// Initialize logger
		logger, err := log.NewLogger(log.Config{
			Type:       "file",
			Level:      "debug",
			JSONFormat: isProduction,
			Filename:   "storage/logs/app.log",
			MaxSize:    100, // 100MB
			MaxBackups: 7,
			MaxAge:     30, // 30 days
			Compress:   true,
		})
		if err != nil {
			return fmt.Errorf("failed to initialize logger: %w", err)
		}

		// Create router
		router := gin.Default()

		router.Use(middleware.Recovery())
		router.Use(requestid.New())
		router.Use(middleware.Security())
		router.Use(middleware.CORS())

		database, err := db.Make()
		if err != nil {
			return err
		}

		// Setup routes
		api.SetupRouter(router)

		// Start server
		srv := &http.Server{
			Addr:    appConfig.Api().Address(),
			Handler: router,
		}

		// Run server in a goroutine
		go func() {
			logger.Info(fmt.Sprintf("Server is running on %s", srv.Addr))
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatal(fmt.Sprintf("Server error: %v", err))
			}
		}()

		util.WaitOSSignalGracefulShutdown(func(ctx context.Context) {
			logger.Info("Shutting down server...")

			if err := srv.Shutdown(ctx); err != nil {
				logger.Error(fmt.Sprintf("Server forced to shutdown: %v", err))
			}

			// Close database connection if exists
			if sqlDB, err := database.DB(); err == nil {
				if err := sqlDB.Close(); err != nil {
					logger.Error(fmt.Sprintf("Error closing database connection: %v", err))
				}
				logger.Info("Database connection closed")
			}

			logger.Info("Server exited properly")
		}, 5*time.Second)

		return nil
	},
}

func init() {
	apiCmd.Flags().StringP("listen", "l", "", "API server listen address. Default in config.yml")
	rootCmd.AddCommand(apiCmd)
}
