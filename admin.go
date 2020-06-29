package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func adminDashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		htmldir+"header.html",
		htmldir+"admin.html",
		htmldir+"footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, "error template read: %v", err)
	}
	err = t.ExecuteTemplate(w, "header", "Admin - Shopsite")
	if err != nil {
		fmt.Fprintf(w, "Error exec header: %v", err)
	}
	err = t.ExecuteTemplate(w, "admin", nil)
	if err != nil {
		fmt.Fprintf(w, "Error exec admin: %v", err)
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Error exec header: %v", err)
	}
}
func adminProducts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		htmldir+"header.html",
		htmldir+"admin.html",
		htmldir+"productlist.html",
		htmldir+"footer.html",
	)
	data := *productsList
	if err != nil {
		fmt.Fprintf(w, "error template read: %v", err)
	}
	err = t.ExecuteTemplate(w, "header", "Admin - Shop")
	if err != nil {
		fmt.Fprintf(w, "Error exec header: %v", err)
	}
	err = t.ExecuteTemplate(w, "productlist", data)
	if err != nil {
		fmt.Fprintf(w, "Error exec product list: %v", err)
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Error exec header: %v", err)
	}
}
