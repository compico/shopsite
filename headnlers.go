package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//
var htmldir string = "./public/html/"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmldir+"index.html", htmldir+"header.html", htmldir+"footer.html")
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "header", "Index - Shop")
	if err != nil {
		fmt.Fprintf(w, "Exec header error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec index error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec footer error: %v", err.Error())
	}
}
func productsHandler(w http.ResponseWriter, r *http.Request) {

	data := getProducts()
	t, err := template.ParseFiles(htmldir+"header.html", htmldir+"products.html", htmldir+"footer.html")
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "header", "Products - Shop")
	if err != nil {
		fmt.Fprintf(w, "Exec header error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "products", data)
	if err != nil {
		fmt.Fprintf(w, "Exec products error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec footer error: %v", err.Error())
	}

}
