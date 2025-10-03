package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
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
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	/** create category and product */
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	product := Product{Name: "Macbook", Price: 20000.00, CategoryID: 1}
	db.Create(&product)

	db.Create(&SerialNumber{Number: "123456789", ProductID: product.ID})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, c := range categories {
		fmt.Println(c.Name, len(c.Products))

		for _, p := range c.Products {
			fmt.Println(p.Name, p.Category.Name, "Serial Number:", p.SerialNumber.Number)
		}
	}

	fmt.Println("--------------------------------")
}
