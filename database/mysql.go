package database

import (
	"fmt"
	"log"

	"github.com/ALTA-BE14-helmimuzkr/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnectionMysql(c *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUSER, c.DBPASS, c.DBHOST, c.DBPORT, c.DBNAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Koneksi database mysql gagal", err.Error())
		return nil
	}

	return db
}
