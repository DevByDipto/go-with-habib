package user

import (
	"ecommerce/config"
	"ecommerce/database"
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
	var reqLogin ReqLogin
	
	// রিকোয়েস্ট বডি থেকে লগইন ডাটা ডিকোড করা
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	// ডাটাবেস থেকে ইউজার খোঁজা
	usr := database.Find(reqLogin.Email, reqLogin.Password)
	
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
    Sub:       usr.ID,
    FirstName: usr.FirstName,
    LastName:  usr.LastName,
})

if err != nil {
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
}
	// লগইন সফল হলে ইউজারের ডাটা পাঠানো
	util.SendData(w, accessToken, http.StatusOK)
}