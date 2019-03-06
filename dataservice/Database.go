package dataservice

import (
	"fmt"
	"log"
	"projects/blog/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Type is the type of database from a Type* constant
type Type string

const (
	// TypeBolt is BoltDB
	TypeBolt Type = "Bolt"
	// TypePostgreSQL is PostgreSQL
	TypePostgreSQL Type = "PostgreSQL"
	// TypeMySQL is MySQL
	TypeMySQL Type = "MySQL"
)

// Info contains the database configurations
type Info struct {
	// database type
	Type Type
	// Mysql Info if used
	MYSQL MySQLInfo
	// Bolt info if used
	Bolt BoltInfo
	// MongoDB info if used
	MongoDB MongoDBInfo
}

// MySQLInfo is the details for the database connection
type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// MongoDBInfo is the details for the database connection
type MongoDBInfo struct {
	URL      string
	Database string
}

// BoltInfo is the details for the database connection
type BoltInfo struct {
	Path string
}

// DSN returns the Data Source Name
func DSN(ci MySQLInfo) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		fmt.Sprintf("%d", ci.Port) +
		")/" +
		ci.Name + "?" + ci.Parameter
}

var dbcon *gorm.DB

/**
*
*
**/
func GetDBConection() *gorm.DB {
	return dbcon
}

/**
* Connect DB
* Run migrate
**/
func InitDb(db Info) {
	var err error
	// Store the config
	switch db.Type {
	case TypeMySQL:
		// Connect MYSQL
		dbcon, err = gorm.Open("mysql", DSN(db.MYSQL))
		if err != nil {
			log.Println("Cannot connect DB MYSQL", err)
		}
		// Migrate the schema
		dbcon.AutoMigrate(&models.User{})
	// case TypePostgreSQL:
	// 	dbcon, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	// 	if err != nil {
	// 		log.Println("Cannot connect DB PostgreSQL", err)
	// 	}
	// 	defer dbcon.Close()
	default:
		log.Println("No registered database in config")
	}

}
