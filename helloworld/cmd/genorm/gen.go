package genorm

import (
	"helloworld/internal/pkg/store"

	"github.com/spf13/cobra"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// GenORMCmd ...
var GenORMCmd = &cobra.Command{
	Use:          "genorm",
	Short:        "Generate ORM code",
	Example:      "api genorm",
	SilenceUsage: true,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		// parse
		return nil
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		g := gen.NewGenerator(gen.Config{
			OutPath:       "./dal/query",
			FieldSignable: true,
			WithUnitTest:  true,
			Mode:          gen.WithQueryInterface,
		})

		var (
			db  *gorm.DB
			err error
		)
		if db, _, err = store.NewMySQL(store.Config{
			//DSN: "root:Xb462415@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local",
			DSN: "root:root@(127.0.0.1:3306)/helloworld?charset=utf8mb4&parseTime=True&loc=Local",
		}); err != nil {
			return err
		}
		g.UseDB(db)

		// g.GenerateModel(tableName string, opts ...gen.ModelOpt)
		g.GenerateAllTable()
		g.ApplyBasic(
			g.GenerateModel("user",
				gen.FieldType("birthday", "sql.NullTime"),
			),
		)

		g.Execute()

		return nil
	},
}
