package database

import(
	. "../models"
	"../config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var err error
var settings config.Configuration

func Open() *gorm.DB {	
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", 
		settings.Server, settings.User, settings.Pass, settings.Port, settings.Database)
	db, _ := gorm.Open("mssql", connString)

	if err != nil {
		fmt.Println(err)
		defer db.Close()
	}
	return db
}

func Migrate() {
	settings = config.GetConfig()
	db := Open()
	db.AutoMigrate(Person{})
	defer db.Close()
}