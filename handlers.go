package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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
	categoryform := r.FormValue("category")
	var (
		data Products
		err  error
	)
	if categoryform == "" {
		data = *productsList
	}
	if categoryform != "" {
		data, err = productsList.getProductsByCategory(categoryform)
		if err != nil {
			fmt.Fprintf(w, "Get products error: %v", err.Error())
		}
	}
	t, err := template.ParseFiles(htmldir+"header.html", htmldir+"leftmenu.html", htmldir+"products.html", htmldir+"footer.html")
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "header", "Products - Shop")
	if err != nil {
		fmt.Fprintf(w, "Exec header error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "leftmenu", getAllCategorys())
	if err != nil {
		fmt.Fprintf(w, "Exec leftmenu error: %v", err.Error())
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

func productHandler(w http.ResponseWriter, r *http.Request) {
	var (
		id          int
		data        Product
		fine        = true
		haveproduct = true
	)

	t, err := template.ParseFiles(
		htmldir+"header.html",
		htmldir+"product.html",
		htmldir+"footer.html",
		htmldir+"error.html",
	)
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
			fmt.Fprintf(w, "Exec errror tempalte error: %v", err.Error())
		}
		fine = false
		haveproduct = false
	}

	if product != "" {
		id, err = strconv.Atoi(product)
		if err != nil {
			err = t.ExecuteTemplate(w, "error", "Conv error:"+err.Error())
			if err != nil {
				fmt.Fprintf(w, "Exec errror tempalte error: %v", err.Error())
			}
			fine = false
		}
	}
	if haveproduct {
		data, err = productsList.getProductById(id)
		if err != nil {
			err = t.ExecuteTemplate(w, "error", err)
			if err != nil {
				fmt.Fprintf(w, "Exec errror tempalte error: %v", err.Error())
			}
			fine = false
		}
		productsList.Product[id].addView()
	}
	if data.IsDeleted {
		err = t.ExecuteTemplate(w, "error", "Товар удалён!")
		if err != nil {
			fmt.Fprintf(w, "Exec errror tempalte error: %v", err.Error())
		}
		fine = false
	}
	if fine {
		err = t.ExecuteTemplate(w, "product", data)
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
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Fprintf(w, "ParseMultipartForm error: %v", err.Error())
		return
	}
	var (
		name        = r.PostFormValue("name")
		description = r.PostFormValue("description")
		category    = r.PostFormValue("category")
		priceval    = r.PostFormValue("price")
	)
	priceval = strings.ReplaceAll(priceval, ",", ".")
	if name == "" || description == "" ||
		category == "" {
		fmt.Fprintln(w, "Error to add: values is empty")
		return
	}
	price, err := strconv.ParseFloat(priceval, 64)
	if err != nil {
		fmt.Fprintf(w, "Error to add, because wrong price: %v", err.Error())
		return
	}
	var p = Product{
		Name:        name,
		Price:       price,
		Description: description,
		Category:    category,
		IsDeleted:   false,
	}
	for _, v := range r.MultipartForm.File {
		for _, v := range v {
			name, err := fileupload(name, globalid, *v)
			if err != nil {
				fmt.Fprintf(w, "Exec footer error: %v", err.Error())
			}
			p.Images = append(p.Images, name)
		}
	}
	productsList.addProduct(p)
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func addReviewHandler(w http.ResponseWriter, r *http.Request) {
	var (
		sID           = r.FormValue("id")
		author        = r.FormValue("author")
		vertues       = r.FormValue("vertues")
		disadvantages = r.FormValue("disadvantages")
		reviewtext    = r.FormValue("reviewtext")
		stars         = r.FormValue("stars")
	)
	id, err := strconv.Atoi(sID)
	if err != nil {
		fmt.Fprintf(w, "Error atoi:%v<br>%v", err,
			"Внутренняя ошибка или кто то пытался в id ввести текст а не число.")
	}
	if len(productsList.Product)-1 < id {
		fmt.Fprintf(w, "Error: %v", "такого продукта нет.")
	}
	if sID == "" || author == "" || reviewtext == "" || stars == "" {
		fmt.Fprintf(w, "Error empty values:%v",
			`ID или автор или текст отзыва или количество звёзд - не указано. <br> Попробуйте заново добавить отзыв`)
	}
	err = productsList.Product[id].addReview(author, vertues,
		disadvantages, reviewtext, stars)
	if err != nil {
		fmt.Fprintf(w, "Err addReview method: %v<br>%v", err,
			"Внутренняя ошибка или кто то пытался в stars отправить направильное число.")
	}
	http.Redirect(w, r, "/product?product="+sID, http.StatusFound)
}

func addtestproducts(w http.ResponseWriter, r *http.Request) {
	testproducts := []Product{
		{
			Images:      []string{"/public/image/testimage/1.jpg"},
			Name:        "Пастельный карандаш",
			Category:    "Пастельные карандаши",
			Description: "Карандаш, пастельный, коричневого цвета",
			Price:       5.428,
		},
		{
			Images:      []string{"/public/image/testimage/2.jpg"},
			Name:        "Графитовый карандаш",
			Category:    "Графитовые карандаши",
			Description: "Графитовый карандаш, твёрдый",
			Price:       7.228,
		},
		{
			Images:      []string{"/public/image/testimage/3.jpg"},
			Name:        "Графитовый карандаш",
			Category:    "Графитовые карандаши",
			Description: "Графитовый карандаш, мягкий",
			Price:       2.7,
		},
		{
			Images:      []string{"/public/image/testimage/4.jpg"},
			Name:        "Набор цветных карандашей",
			Category:    "Восковые карандаши",
			Description: "Цветные карандаши, с восковым ядром. Набор 10 шт.",
			Price:       54.2,
		},
		{
			Images:      []string{"/public/image/testimage/5.jpg"},
			Name:        "Набор цветных карандашей",
			Category:    "Восковые карандаши",
			Description: "Цветные карандаши, с восковым ядром. Набор 16 шт.",
			Price:       80,
		},
		{
			Images:      []string{"/public/image/testimage/6.jpg"},
			Name:        "Набор графитовых карандашей",
			Category:    "Графитовые карандаши",
			Description: "Набор графитовых карандашей. Разной жёсткости. Набор 10 шт.",
			Price:       24.523,
		},
	}
	for i := 0; i < len(testproducts); i++ {
		productsList.addProduct(testproducts[i])
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
