package repository

import (
	"fmt"
	"github.com/mahdi-cpp/go-english/config"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Con() {
	DB = config.DB
}

type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
	DeletedAt gorm.DeletedAt
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  int
	AddressID int
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func OverrideJoinTable() {

	Con()

	DB.Migrator().DropTable(&Person{}, &Address{}, &PersonAddress{})

	if err := DB.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{}); err != nil {
		fmt.Printf("Failed to setup join table for person, got error %v", err)
	}

	if err := DB.AutoMigrate(&Person{}, &Address{}); err != nil {
		fmt.Printf("Failed to migrate, got %v", err)
	}

	address1 := Address{Name: "Mahdi address 1"}
	address2 := Address{Name: "Mahdi address 2"}
	address3 := Address{Name: "Mahdi address 3"}
	person := Person{Name: "Mahdi", Addresses: []Address{address1, address2, address3}}
	DB.Create(&person)

	var addresses1 []Address
	if err := DB.Model(&person).Association("Addresses").Find(&addresses1); err != nil || len(addresses1) != 2 {
		fmt.Printf("Failed to find address, got error %v, length: %v", err, len(addresses1))
	}

	if err := DB.Model(&person).Association("Addresses").Delete(&person.Addresses[0]); err != nil {
		fmt.Printf("Failed to delete address, got error %v", err)
	}

	if len(person.Addresses) != 1 {
		fmt.Printf("Should have one address left")
	}

	if DB.Find(&[]PersonAddress{}, "person_id = ?", person.ID).RowsAffected != 1 {
		fmt.Printf("Should found one address")
	}

	var addresses2 []Address
	if err := DB.Model(&person).Association("Addresses").Find(&addresses2); err != nil || len(addresses2) != 1 {
		fmt.Printf("Failed to find address, got error %v, length: %v", err, len(addresses2))
	}

	if DB.Model(&person).Association("Addresses").Count() != 1 {
		fmt.Printf("Should found one address")
	}

	var addresses3 []Address
	if err := DB.Unscoped().Model(&person).Association("Addresses").Find(&addresses3); err != nil || len(addresses3) != 2 {
		fmt.Printf("Failed to find address, got error %v, length: %v", err, len(addresses3))
	}

	if DB.Unscoped().Find(&[]PersonAddress{}, "person_id = ?", person.ID).RowsAffected != 2 {
		fmt.Printf("Should found soft deleted addresses with unscoped")
	}

	if DB.Unscoped().Model(&person).Association("Addresses").Count() != 2 {
		fmt.Printf("Should found soft deleted addresses with unscoped")
	}

	DB.Model(&person).Association("Addresses").Clear()

	if DB.Model(&person).Association("Addresses").Count() != 0 {
		fmt.Printf("Should deleted all addresses")
	}

	if DB.Unscoped().Model(&person).Association("Addresses").Count() != 2 {
		fmt.Printf("Should found soft deleted addresses with unscoped")
	}

	DB.Unscoped().Model(&person).Association("Addresses").Clear()

	if DB.Unscoped().Model(&person).Association("Addresses").Count() != 0 {
		fmt.Printf("address should be deleted when clear with unscoped")
	}

	//-----------------------------------------------
	address21 := Address{Name: "Ali address 1"}
	address22 := Address{Name: "Ali address 2"}
	person2 := Person{Name: "Ali", Addresses: []Address{address21, address22}}
	DB.Create(&person2)

	//if err := DB.Select(clause.Associations).Delete(&person2).Error; err != nil {
	//	fmt.Printf("failed to delete person, got error: %v", err)
	//}
	//
	//if count := DB.Unscoped().Model(&person2).Association("Addresses").Count(); count != 2 {
	//	fmt.Printf("person's addresses expects 2, got %v", count)
	//}
	//
	//if count := DB.Model(&person2).Association("Addresses").Count(); count != 0 {
	//	fmt.Printf("person's addresses expects 2, got %v", count)
	//}
}
