package main

import (
	"flag"
	"log"

	"github.com/andreylm/basic-go-server.git/pkg/server"

	"github.com/andreylm/basic-go-server.git/pkg/server/config"

	"github.com/andreylm/basic-go-server.git/pkg/db"
)

var serverConfiguration config.ServerConfigurations

func init() {
	var err error
	flag.StringVar(&serverConfiguration.Port, "port", "8000", "Assigning the port")

	flag.StringVar(&serverConfiguration.Driver, "db-driver", "mysql", "Assigning the driver")
	flag.StringVar(&serverConfiguration.DbConfig.Port, "db-port", "3306", "Assigning the db port")
	flag.StringVar(&serverConfiguration.DbConfig.Host, "db-host", "localhost", "Assigning the db host")
	flag.StringVar(&serverConfiguration.DbConfig.Database, "db-database", "test", "Assigning the databse")
	flag.StringVar(&serverConfiguration.DbConfig.User, "db-user", "test", "Assigning the user")
	flag.StringVar(&serverConfiguration.DbConfig.Password, "db-password", "test", "Assigning the password")

	flag.StringVar(&serverConfiguration.KeyPublic, "key-public", "./internal/keys/app.rsa", "Assigning path to public key")
	flag.StringVar(&serverConfiguration.KeyPrivate, "key-private", "./internal/keys/app.rsa.pub", "Assigning path to private key")

	flag.Parse()

	if serverConfiguration.DbConfig.DriverType, err = db.GetDriverType(serverConfiguration.Driver); err != nil {
		log.Panic(err)
	}
}

func main() {
	server := server.NewServer()

	if err := server.Init(&serverConfiguration); err != nil {
		log.Panic(err)
	}

	if err := server.Start(); err != nil {
		log.Panic(err)
	}
}
