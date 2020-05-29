package main

type Products struct {
	Product []Product
}

type Product struct {
	Image       string
	Name        string
	Price       float32
	Description string
	Category    string
	CategoryId  string
}

func getProduct() Product {
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

func getProducts() Products {
	var products Products
	products.Product = make([]Product, 15)

	for i := 0; i < len(products.Product); i++ {
		products.Product[i] = getProduct()
	}
	return products
}
