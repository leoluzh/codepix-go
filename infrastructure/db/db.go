package db

import (
	"github.com/leoluzh/codepix-go/domain/model"

	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	//setting current path
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatal("Error on loading .env file")
	}
}

func ConnectDB(env string) *gorm.DB {

	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv("dsn")
		dbType := os.Getenv("dbType")
		db, err = gorm.Open(dbType, dsn)
	} else {
		dsn = os.Getenv("dnsTest")
		dbTypeTest := os.Getenv("dbTypeTest")
		db, err = gorm.Open(dbTypeTest, dsn)
	}

	if err != nil {
		log.Fatalf("Erro on connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db

}
