package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin

	// Decode the request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	// Find the user in the repository by email and password
	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Create the JWT access token using user details
	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID,       
		FirstName: usr.FirstName,
		LastName:  usr.LastName, 
		Email:     usr.Email,    
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the access token to the client
	util.SendData(w, http.StatusCreated, accessToken)
}