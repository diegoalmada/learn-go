package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     category
	SerialNumber serialNumber
	gorm.Model
}

type serialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&product{}, &category{}, &serialNumber{})

	/** create category and product */
	category := category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&product{Name: "Mouse", Price: 120.00, CategoryID: 1})

	db.Create(&serialNumber{Number: "123456789", ProductID: 1})

	var products []product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}
}
