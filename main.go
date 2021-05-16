package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	port = ":8080"
)

func init() {
	fmt.Println("Starting server http://localhost" + port + "/")
}

func main() {
	server := &http.Server{
		Addr:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/addproduct", addproductHandler)
	http.HandleFunc("/api/addproduct", addproductMethod)
	http.HandleFunc("/api/addreview", addReviewHandler)
	http.HandleFunc("/api/addtestproducts", addtestproducts)
	http.HandleFunc("/admin/products/all", adminProducts)
	http.HandleFunc("/admin", adminDashboard)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v", err.Error())
		return
	}
}
