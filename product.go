package main

import (
	"errors"
	"strconv"
)

type (
	Products struct {
		Product []Product `json:"products"`
	}
	Categorys struct {
		Category []Category
	}
	Product struct {
		Images        []string `json:"images"`
		Name          string   `json:"name"`
		NameId        string   `json:"nameid"`
		Price         float64  `json:"price"`
		Description   string   `json:"description"`
		Category      string   `json:"category"`
		CategoryId    string   `json:"categoryid"`
		Reviews       []Review `json:"reviews"`
		Views         int      `json:"views"`
		AverageRating float64  `json:"averagerating"`
		IsDeleted     bool     `json:"isdeleted"`
		ID            int      `json:"id"`
	}
	Review struct {
		Author        string  `json:"author"`
		Virtues       string  `json:"vertues"`
		Disadvantages string  `json:"disadvantages"`
		ReviewText    string  `json:"reviewtext"`
		Stars         float64 `json:"stars"`
	}
	Category struct {
		Id   string
		Name string
	}
)

var globalid = 0
var hashmap = make(map[string]int)
var categoryids = make(map[string][]int)
var categorname = make(map[string]string)

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
	categoryids[p.CategoryId] = append(categoryids[p.CategoryId], p.ID)
	categorname[p.CategoryId] = p.Category
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

func (products *Products) getProductsByCategory(categoryid string) (Products, error) {
	var p Products
	p.Product = make([]Product, 0)
	for i := 0; i < len(categoryids[categoryid]); i++ {
		x, err := productsList.getProductById(categoryids[categoryid][i])
		if err != nil {
			return p, err
		}
		p.Product = append(p.Product, x)
	}
	return p, nil
}

func (products *Products) getMultipleItems(count int) Products {
	if x := len(products.Product); x < count {
		difference := (count - x) - 1
		count -= difference
	}
	result := Products{
		products.Product[:count],
	}
	return result
}

func (product *Product) addView() {
	product.Views++
}

func (product *Product) addReview(author, vertues, disadvantages, reviewtext, stars string) (err error) {
	r := Review{}
	r.Stars, err = strconv.ParseFloat(stars, 64)
	if err != nil {
		return err
	}
	if r.Stars > 5 || r.Stars <= 0 {
		return errors.New("Неправильное количество звёзд.")
	}
	r.Author = author
	r.Virtues = vertues
	r.Disadvantages = disadvantages
	r.ReviewText = reviewtext
	product.Reviews = append(product.Reviews, r)
	return nil
}

func getAllCategorys() Categorys {
	var x Categorys
	x.Category = make([]Category, 0)
	for k, v := range categorname {
		t := Category{
			Id:   k,
			Name: v,
		}
		x.Category = append(x.Category, t)
	}
	return x
}
