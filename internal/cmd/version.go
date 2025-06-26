package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Version: ", os.Getenv("APP_VERSION"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
