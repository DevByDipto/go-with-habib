package handlers

import (
	"ecommerce/util"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, "this is test", 200)
}