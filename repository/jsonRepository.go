package repository

import (
	"github.com/mahdi-cpp/go-english/config"
	"github.com/mahdi-cpp/go-english/models"
)

func ConnectDatabase() {
	config.Database()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}

func CreatNewUser() {
	//models.CreatUsers()
	models.QueryUsers()
}

//type UserWithJSON struct {
//	gorm.Model
//	Name       string
//	Attributes datatypes.JSON
//}
//
//func JsonQuery() {
//
//	config.DB.Create(&UserWithJSON{
//		Name:       "json-1",
//		Attributes: datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
//	}
//
//	// Check JSON has keys
//	datatypes.JSONQuery("attributes").HasKey(value, keys...)
//
//	db.Find(&user, datatypes.JSONQuery("attributes").HasKey("role"))
//	db.Find(&user, datatypes.JSONQuery("attributes").HasKey("orgs", "orga"))
//	// MySQL
//
//}
