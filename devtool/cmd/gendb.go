package cmd

import (
	"app/db"

	"github.com/spf13/cobra"
	"gorm.io/gen"
)

var GendbCmd = &cobra.Command{
	Use:   "gen:db",
	Short: "Generate database schema",
	Long:  `Generate database schema.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gen.NewGenerator(gen.Config{
			OutPath:           "db",
			Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
			OutFile:           "gen.go",
			ModelPkgPath:      "models",
			WithUnitTest:      false,
			FieldNullable:     false,
			FieldCoverable:    false,
			FieldSignable:     false,
			FieldWithIndexTag: false,
			FieldWithTypeTag:  false,
		})

		db, err := db.Make()
		if err != nil {
			return err
		}

		g.UseDB(db)

		userTable := g.GenerateModel("users", gen.FieldType("extra", "datatypes.JSON"), gen.FieldType("status", "int8"))

		g.ApplyBasic(
			userTable,
		)

		g.Execute()
		return nil
	},
}
