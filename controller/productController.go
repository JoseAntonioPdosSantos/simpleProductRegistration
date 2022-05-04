package controller

import (
	"log"
	"net/http"
	"simpleProductRegistration/model"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := model.FindAll()
	templates.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error on convert price:", err)
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error on convert quantity:", err)
		}

		model.Create(name, description, priceConvert, quantityConvert)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	model.Remove(idProduct)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := model.FindOneById(idProduct)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error on convert price:", err)
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error on convert quantity:", err)
		}

		model.Update(id, name, description, priceConvert, quantityConvert)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}
