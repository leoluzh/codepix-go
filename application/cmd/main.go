package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/leoluzh/codepix-go/application/grpc"
	"github.com/leoluzh/codepix-go/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
