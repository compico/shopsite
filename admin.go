package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/compico/shopsite/internal/dataworker"
)

func adminDashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		htmldir+"header.html",
		htmldir+"admin.html",
		htmldir+"dashboard.html",
		htmldir+"footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, "error template read: %v", err)
	}
	data := datah.GetDataAndChangeTitle("Admin")
	err = t.ExecuteTemplate(w, "dashboard", data)
	if err != nil {
		fmt.Fprintf(w, "Error exec admin: %v", err)
	}
}
func adminProducts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		htmldir+"header.html",
		htmldir+"admin.html",
		htmldir+"productlist.html",
		htmldir+"footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, "error template read: %v", err)
	}
	data := datah.GetDataAndChangeTitle("Admin")
	data.Data = *dataworker.ProductsList
	err = t.ExecuteTemplate(w, "productlist", data)
	if err != nil {
		fmt.Fprintf(w, "Error exec product list: %v", err)
	}
}
