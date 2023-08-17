package main

import (
	"fmt"
	"strings"

	"gorm.io/gen"
	"gorm.io/gorm"
	_ "gorm.io/plugin/soft_delete"
)

const (
	MYSQLDSN = "root:root@tcp(localhost:3306)/jiaqin?charset=utf8mb4&parseTime=True"
	SQLPATH  = "../sql/"
)

func init() {
	DB = ConnectDB(MYSQLDSN)
	if err := ImportSQL(DB, SQLPATH); err != nil {
		panic(fmt.Errorf("ImportSQL failed:%w", err))
	}
}

var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	"int":    func(columnType gorm.ColumnType) (dataType string) { return "int32" },
	"bigint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	"json":   func(columnType gorm.ColumnType) string { return "json.RawMessage" },
	// bool mapping
	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasPrefix(ct, "tinyint(1)") {
			return "bool"
		}
		return "int"
	},
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../dal/query",
		ModelPkgPath: "../dal/model",
		WithUnitTest: true,
		// generate model global configuration
		FieldNullable: true, // generate pointer when field is nullable
		//FieldCoverable:    true, // generate pointer when field has default value
		FieldWithIndexTag: true, // generate with gorm index tag
		FieldWithTypeTag:  true, // generate with gorm column type tag
	})
	g.UseDB(DB)

	// specify diy mapping relationship
	g.WithDataTypeMap(dataMap)

	// generate all field with json tag end with "_example"
	// g.WithJSONTagNameStrategy(func(c string) string { return c + "_example" })
	g.WithJSONTagNameStrategy(func(c string) string { return "-" })

	holidaySettingsTable := g.GenerateModel("holiday_settings",
		gen.FieldType("deleted_at", "soft_delete.DeletedAt"),
	)

	holidayAffairsSettingsTable := g.GenerateModel("holiday_affairs_settings",
		gen.FieldType("deleted_at", "soft_delete.DeletedAt"),
	)

	holidayAffairsTable := g.GenerateModel("holiday_affairs",
	 //gen.FieldType("deleted_at", "soft_delete.DeletedAt"),
	)

	g.ApplyBasic(
		holidaySettingsTable,
		holidayAffairsSettingsTable,
		holidayAffairsTable,
	)
	// g.ApplyBasic(g.GenerateAllTable()...) // generate all table in db server

	g.Execute()
}
