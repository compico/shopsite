package dataworker

import (
	"errors"
	"strconv"

	"github.com/compico/shopsite/internal/utils"
)

func InitProducts() *Products {
	products := new(Products)
	products.Product = make([]Product, 0)
	return products
}

func (products *Products) AddProduct(p Product) {
	p.ID = Globalid
	p.NameId = strconv.Itoa(Globalid) + "_" + utils.Transcript(p.Name)
	p.CategoryId = utils.Transcript(p.Category)
	p.Price = utils.Round(p.Price, 0.05)
	Categoryids[p.CategoryId] = append(Categoryids[p.CategoryId], p.ID)
	Categorname[p.CategoryId] = p.Category
	Hashmap[p.NameId] = p.ID
	Globalid++
	products.Product = append(products.Product, p)
}

func (products *Products) DelProduct(id int) {
	products.Product[id].IsDeleted = true
}

func (products *Products) EditProduct(id int, p Product) {
	p.ID = id
	products.Product[id] = p
}

func (products *Products) GetProductById(id int) (x Product, err error) {
	if id > (len(products.Product) - 1) {
		err = errors.New("Товар не найден!")
		return x, err
	}
	x = products.Product[id]
	return x, nil
}

func (products *Products) GetProductsByCategory(categoryid string) (Products, error) {
	var p Products
	p.Product = make([]Product, 0)
	for i := 0; i < len(Categoryids[categoryid]); i++ {
		x, err := ProductsList.GetProductById(Categoryids[categoryid][i])
		if err != nil {
			return p, err
		}
		p.Product = append(p.Product, x)
	}
	return p, nil
}

func (products *Products) GetMultipleItems(count int) Products {
	if x := len(products.Product); x < count {
		difference := (count - x) - 1
		count -= difference
	}
	result := Products{
		products.Product[:count],
	}
	return result
}

func (product *Product) AddView() {
	product.Views++
}

func (product *Product) AddReview(author, vertues, disadvantages, reviewtext, stars string) (err error) {
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
