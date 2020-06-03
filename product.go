package main

import (
	"errors"
	"strconv"
)

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
		NameId      string  `json:"nameid"`
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
	Category struct {
		RuName string
		IDs    []int
	}
	Categorys map[string]Category
)

var globalid = 0
var hashmap = make(map[string]int)

func initProducts() *Products {
	products := new(Products)
	products.Product = make([]Product, 0)
	return products
}

func (products *Products) addProduct(p Product) {
	p.ID = globalid
	p.NameId = strconv.Itoa(globalid) + "_" + transcript(p.Name)
	p.CategoryId = transcript(p.Category)
	p.Price = round(p.Price, 0.05)
	hashmap[p.NameId] = p.ID
	globalid++
	products.Product = append(products.Product, p)
}

func (products *Products) delProduct(id int) {
	products.Product[id].IsDeleted = true
}

func (products *Products) editProduct(id int, p Product) {
	p.ID = id
	products.Product[id] = p
}

func (products *Products) getProductById(id int) (x Product, err error) {
	if id > len(products.Product)-1 {
		err = errors.New("Товар не найден!")
		return x, err
	}
	x = products.Product[id]
	return x, nil
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
