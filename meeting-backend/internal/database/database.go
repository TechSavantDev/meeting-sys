package database

import (
	"backend/internal/config"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db Database

func DB() Database {
	if db == nil {
		db = New()
	}
	return db
}

type Database interface {
	Close() error
}

type databaseImpl struct {
	db *gorm.DB
}

func New() Database {
	return newDatabase()
}

func newDatabase() Database {
	db := &databaseImpl{}
	db.initialize()
	return db
}

func (d *databaseImpl) initialize() {
	var err error
	d.db, err = gorm.Open(mysql.Open(
		config.Config().DatabaseDSN()),
		&gorm.Config{},
	)
	if err != nil {
		fmt.Println("连接数据库失败")
		os.Exit(1)
	}
	fmt.Println("连接数据库成功")
}

func (d *databaseImpl) Close() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
