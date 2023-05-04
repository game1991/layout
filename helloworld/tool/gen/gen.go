package main

import (
	"flag"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var dsn string
var sqlType int // 1 mysql 2 pgsql
var outPath string

/*
mysql eg:
"root:root@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
pg eg:
"host=localhost user=root password=root dbname=demo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
*/

func init() {
	flag.StringVar(&dsn, "dsn", "root:root@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local", "gorm解析dsn;例如 root:root@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	flag.IntVar(&sqlType, "type", 1, "数据库类型:1、mysql 2、pgsql")
	flag.StringVar(&outPath, "o", "../../dal/query", "代码生成地址,可以使用相对路径")
}

func main() {

	flag.Parse()

	var gHandler GenHandler
	switch sqlType {
	case 1:
		gHandler = MySQL
	case 2:
		gHandler = PgSQL
	}

	g := gHandler(dsn, outPath)

	// Generate the code
	g.Execute()
}

type GenHandler func(dsn, outPath string) *gen.Generator

func MySQL(dsn, outPath string) *gen.Generator {
	g := gen.NewGenerator(gen.Config{
		OutPath:       outPath,
		FieldSignable: true,
		WithUnitTest:  true,
		Mode:          gen.WithQueryInterface,
	})

	gormdb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	g.UseDB(gormdb) // reuse your gorm db
	// g.GenerateModel(tableName string, opts ...gen.ModelOpt)
	g.GenerateAllTable()
	g.ApplyBasic(
		g.GenerateModel("user",
			gen.FieldType("birthday", "sql.NullTime"),
		),
		g.GenerateModel("account",
			gen.FieldType("group_id", "uint32"),
			gen.FieldType("is_super_user", "bool"),
			gen.FieldType("is_admin", "bool"),
		),
	)
	return g
}

func PgSQL(dsn, outPath string) *gen.Generator {
	g := gen.NewGenerator(gen.Config{
		OutPath: outPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	gormdb, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	g.UseDB(gormdb) // reuse your gorm db

	g.ApplyBasic(
		g.GenerateModel("user_user", gen.FieldType("id", "uint32"), gen.FieldType("group_id", "uint32"), gen.FieldType("zonst_id", "uint32")),
		g.GenerateModel("job_job"),
		g.GenerateModel("group_group", gen.FieldType("id", "uint32")),
		g.GenerateModel("group_owner", gen.FieldType("id", "uint32"), gen.FieldType("group_id", "uint32"), gen.FieldType("user_id", "uint32")),
	)
	return g
}
