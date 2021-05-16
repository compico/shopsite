package dataworker

import "github.com/compico/shopsite/internal/config"

type (
	Data struct {
		Config    config.Config
		Data      interface{}
		Error     string
		Categorys Categorys
	}
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
