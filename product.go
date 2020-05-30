package main

type (
	Products struct {
		Product []Product
	}

	Reviews struct {
		Review []Review
	}

	Product struct {
		Image       string
		Name        string
		Price       float32
		Description string
		Category    string
		CategoryId  string
		Reviews     Reviews
	}

	Review struct {
		Author        string
		Virtues       string
		Disadvantages string
		ReviewText    string
		Stars         float32
	}
)

func initProducts() *Products {
	products := new(Products)
	products.Product = make([]Product, 0)
	return products
}

func (products *Products) addProduct() {

}

func getTestProduct() Product {
	product := Product{
		Image:       "public/image/products/1.jpg",
		Name:        "Карандаш",
		Price:       6.40,
		Description: "Мягкий графитовый карандаш",
		Category:    "Графитовые",
		CategoryId:  "grafitovie",
	}
	return product
}

func getTestProducts() Products {
	var products Products
	products.Product = make([]Product, 15)

	for i := 0; i < len(products.Product); i++ {
		products.Product[i] = getTestProduct()
	}
	return products
}
