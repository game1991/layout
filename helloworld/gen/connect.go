package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("connect db fail:%w", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("get DB failed:%w", err))
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.Apply(&gorm.Config{
		ConnPool: sqlDB,
	}); err != nil {
		panic(err)
	}

	return db
}

var (
	ErrImportSQLFailure     = errors.New("数据库导入失败")
	ErrSqlPathNotFound      = errors.New("数据库SQL文件不存在")
	ErrSqlPathReadFailed    = errors.New("数据库SQL文件读取失败")
	ErrImportSQLEmptyParams = errors.New("ImportSQL请求参数为空，请检查入参")
)

func ImportSQL(db *gorm.DB, sqlPath string) error {
	if db == nil || sqlPath == "" {
		return ErrImportSQLEmptyParams
	}
	s, err := os.Stat(sqlPath)
	if os.IsNotExist(err) {
		return errors.Join(ErrSqlPathNotFound, err)
	}

	var sqls []byte
	var sqlArr []string

	if s.IsDir() {
		// 递归目录下所有文件
		err = filepath.Walk(sqlPath, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				fileContent, err := os.ReadFile(path)
				if err != nil {
					log.Println("递归目录文件读取失败：" + err.Error())
					return err
				}
				ext := filepath.Ext(path)
				if ext == ".sql" {
					sqlArr = append(sqlArr, strings.Split(string(fileContent), ";")...)
				}
			}
			return nil
		})
		if err != nil {
			return errors.Join(ErrSqlPathReadFailed, err)
		}

	} else {
		// 指定文件路径读取
		if !strings.HasSuffix(sqlPath, ".sql") {
			return errors.Join(ErrSqlPathReadFailed, fmt.Errorf("不支持的文件类型[%s]，仅支持.sql文件", filepath.Ext(sqlPath)))
		}
		sqls, err = os.ReadFile(sqlPath)
		if err != nil {
			return errors.Join(ErrSqlPathReadFailed, err)
		}
		sqlArr = strings.Split(string(sqls), ";")
	}

	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := db.Exec(sql).Error
		if err != nil {
			return errors.Join(ErrImportSQLFailure, err)
		}
		log.Println(sql, "\t success!")
	}
	return nil
}
