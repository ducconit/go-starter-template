package cmd

import (
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var MigrationCmd = &cobra.Command{
	Use:   "make:migration [name]",
	Short: "Make migration",
	Long:  `Make migration SQL file using Goose with the specified name. Default file type is sql.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		dir, _ := cmd.Flags().GetString("dir")
		fileType, _ := cmd.Flags().GetString("type")
		if fileType != "sql" && fileType != "go" {
			return fmt.Errorf("file type must be sql or go")
		}

		goose.SetSequential(true)

		return goose.Create(nil, dir, name, fileType)
	},
}

func init() {
	MigrationCmd.PersistentFlags().StringP("dir", "d", "migrations", "Migration directory")
	MigrationCmd.PersistentFlags().StringP("type", "t", "sql", "Migration file type (sql or go)")
}
