package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})
	//db.Create(&Product{Name: "Notebook", Price: 4999.00})
	/*products := []Product{
		{Name: "Tablet", Price: 299.00},
		{Name: "Mouse", Price: 19.00},
		{Name: "Headset", Price: 199.00},
		{Name: "Keyboard", Price: 199.00},
	}

	db.Create(&products)*/

	/** find one **/
	/*var product Product
	//db.First(&product, 2)

	db.First(&product, "name = ?", "Headset")
	fmt.Println(product)*/

	/** get all **/
	/*var products []Product
	//db.Limit(2).Find(&products)
	db.Where("price >= ?", 199.1).Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}*/

	/** update product **/
	/*var p Product
	db.First(&p, 8)
	p.Name = "New Mouse"
	db.Save(&p)*/

	/** delete product **/
	var p2 Product
	db.First(&p2, 8)
	db.Delete(&p2)
}
