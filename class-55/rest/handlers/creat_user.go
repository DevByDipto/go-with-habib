package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	
	// r.Body থেকে JSON ডাটা ডিকোড করে newUser ভেরিয়েবলে বসানো
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	// মেথড কল করে ইউজার স্টোর করা
	createdUser := newUser.Store()

	// সফল হলে JSON রেসপন্স পাঠানো
	util.SendData(w, createdUser, http.StatusCreated)
}