package product

import (
	
	"ecommerce/util"
	"net/http"
	"strconv"
)


func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch the list of all products from the repository
	// Get limit and page from query params
	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	// Convert strings to int64; ignoring errors for simplicity as per the image
	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	// Set default values if params are missing or zero
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}
	productList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Fetch total count of items from the service
	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	

	// Send the successful response back to the client
	util.SendPage(w, productList, page,limit,cnt)

	
}
