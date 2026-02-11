package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// URL থেকে আইডি নেওয়া
	productIDStr := r.PathValue("id")

	// স্ট্রিং আইডিকে ইন্টিজারে রূপান্তর
	pId, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w,"Please give me a valid product id", 400)
		return
	}

	database.Delete(pId)

	// সফল হলে ডাটা পাঠানো
	util.SendData(w, "product delete successfull", 200)
}