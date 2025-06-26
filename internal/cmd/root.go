package cmd

import (
	"core/config"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"

	appConfig "app/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Application skeleton",
	Long:  `Application skeleton.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		configPath, _ := cmd.Flags().GetString("config")
		if configPath == "" {
			return fmt.Errorf("config path is required")
		}
		appenv, _ := cmd.Flags().GetString("env")
		if appenv == "" {
			appenv = "development"
		}

		// if config.yml not exist, create it from config.yml.example
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			f, err := os.Create(configPath)
			if err != nil {
				return err
			}
			defer f.Close()
			src, err := os.Open("config.example.yml")
			if err != nil {
				return err
			}
			defer src.Close()
			_, err = io.Copy(f, src)
			if err != nil {
				return err
			}
		}

		cfg, err := config.Load(config.Option{
			// EnvPrefix: "APP",
			FilePath: configPath,
			Env:      appenv,
			LoadEnv:  true,
		})
		if err != nil {
			return err
		}

		cfg.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			// reload app config
			appConfig.Unmarshal()
		})

		// create symlink from storage/uploads/public to public/storage if not exist
		if _, err := os.Stat("public/storage"); os.IsNotExist(err) {
			if err := os.Symlink("storage/uploads/public", "public/storage"); err != nil {
				// is error `A required privilege is not held by the client.` in windows
				if strings.Contains(err.Error(), "A required privilege is not held by the client.") {
					fmt.Println("[WARNING] Symlink not created. Please run with admin permission.")
				}
			}
		}

		return appConfig.Unmarshal()
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.yml", "Path to config file")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	rootCmd.PersistentFlags().StringP("env", "e", env, "Environment (development, staging, production,...)")

	rootCmd.RunE = apiCmd.RunE
}

func Execute() error {
	return rootCmd.Execute()
}
