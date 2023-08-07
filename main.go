package main

import (
	"flag"
	"go-restapi/controllers"
	"go-restapi/database"

	"github.com/gin-gonic/gin"
)

// go run main.go -host=localhost -port=5432 -user=your_username -pass=your_password -dbname=your_database_name -sslmode=disable -timezone=Europe/Istanbul -connect_timeout=5

func main() {
	var settings database.DatabaseSettings

	host := flag.String("host", "localhost", "Host for database conn")
	port := flag.String("port", "5432", "port for db")
	user := flag.String("user", "postgres", "Username for database connection")
	pass := flag.String("pass", "postgres", "Password for database connection")
	dbname := flag.String("dbname", "tasks", "Database name for connection")
	sslmode := flag.String("sslmode", "disable", "SSL mode for database connection")
	timezone := flag.String("timezone", "Default", "Timezone for database connection")
	connectTimeout := flag.String("connect_timeout", "5", "Connect timeout for database connection")
	flag.Parse()

	settings = database.DatabaseSettings{
		Host:              *host,
		Port:              *port,
		User:              *user,
		Dbname:            *dbname,
		Pass:              *pass,
		Sslmode:           *sslmode,
		Timezone:          *timezone,
		ConnectionTimeout: *connectTimeout,
	}
	dsn := database.GenerateDSN(settings)

	r := gin.Default()

	_, err := database.ConnectDatabase(dsn)
	if err != nil {
		panic(err)
	}

	r.GET("/tasks", controllers.GetAllTasks)
	r.Run(":8080")

}
