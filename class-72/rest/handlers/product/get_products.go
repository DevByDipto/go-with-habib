package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch the list of all products from the repository
	productList, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Successfully return the list of products
	util.SendData(w, http.StatusOK, productList)
}