package main

type (
	Products struct {
		Product []Product `json:"products"`
	}
	Reviews struct {
		Review []Review `json:"reviews"`
	}
	Product struct {
		Image       string  `json:"image"`
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		CategoryId  string  `json:"categoryid"`
		Reviews     Reviews `json:"reviews"`
		IsDeleted   bool    `json:"isdeleted"`
		ID          int     `json:"id"`
	}
	Review struct {
		Author        string  `json:"author"`
		Virtues       string  `json:"vertues"`
		Disadvantages string  `json:"disadvantages"`
		ReviewText    string  `json:"reviewtext"`
		Stars         float64 `json:"stars"`
	}
)

var globalid = 0

func initProducts() *Products {
	products := new(Products)
	products.Product = make([]Product, 0)
	return products
}

func (products *Products) addProduct(p Product) {
	p.ID = globalid
	products.Product = append(products.Product, p)
}

func (products *Products) delProduct(id int) {
	products.Product[id].IsDeleted = true
}

func (products *Products) editProduct(id int, p Product) {
	p.ID = id
	products.Product[id] = p
}

func (products *Products) getProductById(id int) Product {
	x := products.Product[id]
	return x
}

func (products *Products) getMultipleItems(count int) Products {
	if x := len(products.Product); x < count {
		count = x - 1
	}
	result := Products{
		products.Product[:count],
	}
	return result
}
