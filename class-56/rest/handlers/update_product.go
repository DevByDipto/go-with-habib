package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// URL থেকে আইডি নেওয়া
	productIDStr := r.PathValue("id")

	// স্ট্রিং আইডিকে ইন্টিজারে রূপান্তর
	pId, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}
	newProduct.ID = pId
	// ডাটাবেস থেকে প্রোডাক্ট খোঁজা
	database.Update(newProduct)

	// সফল হলে ডাটা পাঠানো
	util.SendData(w, "product update successfull", 201)
}
