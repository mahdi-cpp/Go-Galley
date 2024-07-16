package repository

import (
	"fmt"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Info     map[string]interface{} `gorm:"type:json;serializer:json"`
	Persians map[string]interface{} `gorm:"type:json;serializer:json"`
}

func UserStart() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
	//CreatUsers()
	QueryUsers()
}

func CreatUsers() {

	user1 := User{
		Name:  "Mahdi",
		Email: "mahdi.cpp@gmail.com",
		Info: map[string]interface{}{
			"age":    30,
			"city":   "Tehran",
			"height": 170,
			"weight": 70,
		},
		Persians: map[string]interface{}{
			"type1":    "Verb",
			"type2":    "Noun",
			"persian1": []string{"mahdi", "ali", "reza"},
			"persian2": []string{"sara", "maryam", "leyla"},
		},
	}

	user2 := User{
		Name:  "Ali3",
		Email: "ali.1392@gmail.com",
		Info: map[string]interface{}{
			"age":    24,
			"city":   "tabrize",
			"height": 100,
			"weight": 87,
		},
		Persians: map[string]interface{}{
			"type1":    "Verb",
			"type2":    "Noun",
			"persian1": []string{"mahdi", "ali", "reza"},
			"persian2": []string{"sara", "maryam", "leyla"},
		},
	}

	DB.Create(&user1)
	DB.Create(&user2)
}

func QueryUsers() {
	var users []User

	DB.Where("persians->>'type1' = ?", "Adjective").Find(&users)

	fmt.Println("--------------------------")
	for _, u := range users {
		var p = u.Persians["type1"]
		fmt.Println(u.Name)
		fmt.Println(p)
	}
	fmt.Println("--------------------------")
}

//
//func QueryUsers2() {
//	var users []User
//
//	DB.Where("info->>'width' = ?", 11).Find(&users)
//	DB.Where("info->>'width' = ?", 11).Find(&users)
//
//	fmt.Println("query result: ")
//	for _, u := range users {
//		var city = u.Info["city"]
//		fmt.Println(city)
//	}
//
//	fmt.Println("-----------------------")
//}

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
