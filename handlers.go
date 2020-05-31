package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

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
	t, err := template.ParseFiles(htmldir+"header.html", htmldir+"products.html", htmldir+"footer.html")
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "header", "Products - Shop")
	if err != nil {
		fmt.Fprintf(w, "Exec header error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "products", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec products error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec footer error: %v", err.Error())
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		htmldir+"header.html",
		htmldir+"product.html",
		htmldir+"footer.html",
		htmldir+"error.html",
	)
	fine := true
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "header", "Product - Shop")
	if err != nil {
		fmt.Fprintf(w, "Exec header error: %v", err.Error())
	}
	product := r.FormValue("product")
	if product == "" {
		err = t.ExecuteTemplate(w, "error", "Товар не найден!")
		if err != nil {
			fmt.Fprintf(w, "Exec products error: %v", err.Error())
		}
		fine = false
	}
	if fine {
		err = t.ExecuteTemplate(w, "product", nil)
		if err != nil {
			fmt.Fprintf(w, "Exec products error: %v", err.Error())
		}
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec footer error: %v", err.Error())
	}
}

func addproductHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmldir+"header.html", htmldir+"addproduct.html", htmldir+"footer.html")
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "header", "Добавить товар - Shop")
	if err != nil {
		fmt.Fprintf(w, "Exec header error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "addproduct", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec addproduct error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec footer error: %v", err.Error())
	}
}

func addproductMethod(w http.ResponseWriter, r *http.Request) {

	var (
		image       = r.PostFormValue("image")
		name        = r.PostFormValue("name")
		description = r.PostFormValue("description")
		category    = r.PostFormValue("category")
		categoryid  = r.PostFormValue("categoryid")
	)

	if image == "" || name == "" || description == "" ||
		category == "" || categoryid == "" {
		fmt.Fprintln(w, "Error to add: values is empty")
		return
	}

	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		fmt.Fprintf(w, "Error to add, because wrong price: %v", err.Error())
		return
	}
	p := Product{
		Image:       image,
		Name:        name,
		Price:       price,
		Description: description,
		Category:    category,
		CategoryId:  categoryid,
		Reviews:     Reviews{nil},
		IsDeleted:   false,
	}
	productsList.addProduct(p)
	http.Redirect(w, r, "/", http.StatusFound)
}
