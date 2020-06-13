package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	port         = ":8080"
	productsList = initProducts()
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
	http.HandleFunc("/testfile", addtestfilehandler)
	http.HandleFunc("/api/addproduct", addproductMethod)
	http.HandleFunc("/api/addtestproducts", addtestproducts)
	http.HandleFunc("/api/testfile", testfilehandler)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v", err.Error())
		return
	}
}

func addtestfilehandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles(htmldir+"header.html", htmldir+"test.html", htmldir+"footer.html")
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "header", "Products - Shop")
	if err != nil {
		fmt.Fprintf(w, "Exec header error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "test", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec products error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Exec footer error: %v", err.Error())
	}
}

func testfilehandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) //10 MB
	if err != nil {
		fmt.Fprintf(w, "Parsiong Multipart Form error: %v", err.Error())
		return
	}
	files, _, err := r.FormFile("files")
	if err != nil {
		fmt.Fprintf(w, "FormFile error: %v", err.Error())
		return
	}
	defer files.Close()
	tempFile, err := ioutil.TempFile("./public/image/png", "*.png")
	if err != nil {
		fmt.Fprintf(w, "TempFile error: %v", err.Error())
		return
	}
	defer tempFile.Close()
	globalid++
	fileBytes, err := ioutil.ReadAll(files)
	if err != nil {
		fmt.Fprintf(w, "ReadAll error: %v", err.Error())
		return
	}
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		fmt.Fprintf(w, "Write error: %v", err.Error())
		return
	}
}
