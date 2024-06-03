package main

import "github.com/compico/shopsite/internal/dataworker"

func getAllCategorys() dataworker.Categorys {
	var x dataworker.Categorys
	x.Category = make([]dataworker.Category, 0)
	for k, v := range dataworker.CategorName {
		t := dataworker.Category{
			Id:   k,
			Name: v,
		}
		x.Category = append(x.Category, t)
	}
	return x
}
