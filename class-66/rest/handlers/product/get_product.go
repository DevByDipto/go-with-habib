package product

import (
	
	"ecommerce/util"
	"net/http"
	"strconv"
)
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the path parameter
	productID := r.PathValue("id")

	// Convert the string ID to an integer
	pId, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	// Fetch the specific product from the repository
	product, err := h.productRepo.Get(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Handle case where the repository returns no product
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	// Successfully return the product data
	util.SendData(w, http.StatusOK, product)
}