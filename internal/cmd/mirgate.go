package cmd

import (
	"app/db"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [up|down|status|version]",
	Short: "Run migrations",
	Long:  `Run migrations that have not been applied.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			args = []string{"up"}
		}
		arg := args[0]
		dir, _ := cmd.Flags().GetString("dir")

		g, err := db.Make()
		if err != nil {
			return err
		}
		db, err := g.DB()
		if err != nil {
			return err
		}
		return goose.RunContext(cmd.Context(), arg, db, dir, args[1:]...)
	},
}

func init() {
	migrateCmd.Flags().StringP("dir", "d", "migrations", "Migration directory")
	rootCmd.AddCommand(migrateCmd)
}
