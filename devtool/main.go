package main

import (
	"core/config"
	"fmt"
	"os"
	"time"

	"app/devtool/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devtool",
	Short: "Devtool",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		configPath, _ := cmd.Flags().GetString("config")
		if configPath == "" {
			return fmt.Errorf("config path is required")
		}
		appenv, _ := cmd.Flags().GetString("env")
		if appenv == "" {
			appenv = "development"
		}

		_, err := config.Load(config.Option{
			// EnvPrefix: "APP",
			FilePath: configPath,
			Env:      appenv,
			LoadEnv:  true,
		})
		return err
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.yml", "Path to config file")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	rootCmd.PersistentFlags().StringP("env", "e", env, "Environment (development, staging, production,...)")
}

func main() {
	time.Local = time.UTC

	rootCmd.AddCommand(cmd.GendbCmd)
	rootCmd.AddCommand(cmd.MigrationCmd)
	rootCmd.AddCommand(cmd.GenSwaggerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
