package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// URL থেকে আইডি নেওয়া
	productIDStr := r.PathValue("id")

	// স্ট্রিং আইডিকে ইন্টিজারে রূপান্তর
	pId, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	// ডাটাবেস থেকে প্রোডাক্ট খোঁজা
	product := database.Get(pId)
	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	// সফল হলে ডাটা পাঠানো
	util.SendData(w, product, 200)
}