package model

import (
	"simpleProductRegistration/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAll() []Product {
	products := []Product{}

	database := db.ConnectionToDB()
	allProducts, err := database.Query("select * from product order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)

	}

	defer database.Close()

	return products
}

func Create(name string, description string, price float64, quantity int) {
	database := db.ConnectionToDB()

	insert, err := database.Prepare("insert into product (name, description, price, quantity) values ($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)

	defer database.Close()
}

func Remove(id string) {
	database := db.ConnectionToDB()
	removeProduct, err := database.Prepare("delete from product where id=$1")
	if err != nil {
		panic(err.Error())
	}
	removeProduct.Exec(id)
	defer database.Close()
}

func FindOneById(id string) Product {
	database := db.ConnectionToDB()
	editProduct, err := database.Query("select * from product where id=" + id)

	if err != nil {
		panic(err.Error())
	}
	product := Product{}

	for editProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = editProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id
		product.Name = name
		product.Description = description
		product.Quantity = quantity
		product.Price = price
	}
	defer database.Close()
	return product
}

func Update(id string, name string, description string, price float64, quantity int) {
	database := db.ConnectionToDB()

	insert, err := database.Prepare("update product set name=$1, description=$2,price=$3,quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity, id)

	defer database.Close()
}
