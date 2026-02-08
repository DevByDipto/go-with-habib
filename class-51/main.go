// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type Product struct {
// 	ID          int    `json:"id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Price       string `json:"price"`
// 	ImgUrl      string `json:"imgUrl"`
// }

// var productList []Product

// func getProducts(w http.ResponseWriter, r *http.Request) {

// 	sendData(w, productList, 200)
// }

// // createProductHandler
// func createProduct(w http.ResponseWriter, r *http.Request) {

// 	var newProduct Product
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&newProduct)

// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Plz give me valid json", 400)
// 		return
// 	}

// 	newProduct.ID = len(productList) + 1
// 	productList = append(productList, newProduct)

// 	sendData(w, newProduct, 201)
// }

// func handleCors(w http.ResponseWriter) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Method", "GET,POST,PUT,DELETE,OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Habib")
// 	w.Header().Set("Content-Type", "application/json")
// }

// func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader((200))
// 	}
// }

// func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
// 	w.WriteHeader(statusCode)
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(data)
// }

// func main() {
// 	mux := http.NewServeMux()

// 	mux.Handle("GET /products", http.HandlerFunc(getProducts)) // route
// 	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

// 	fmt.Println("server runing 8080")

// 	globalRouter := globalRouter(mux)
// 	err := http.ListenAndServe(":8080", globalRouter)

// 	if err != nil {
// 		fmt.Println("Error starting the server", err)
// 	}

// }

// func init() {
// 	// product1:=
// }

// func globalRouter(mux *http.ServeMux) http.Handler {
// 	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib")
// 			w.Header().Set("Content-Type", "application/json")

// 			if r.Method == "OPTIONS" {

// 			w.WriteHeader(200)
// 			return
// 		} else {
// 			mux.ServeHTTP(w, r)
// 		}
// 	}
// 	return http.HandlerFunc(handleAllReq)
// }

// -------------------------------

package main

import (
	"ecommerce/cmd"
	
)



// func handleCors(w http.ResponseWriter) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Method", "GET,POST,PUT,DELETE,OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Habib")
// 	w.Header().Set("Content-Type", "application/json")
// }

// func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader((200))
// 	}
// }



func main() {
	cmd.Serve()

}

func init() {
	// product1:=
}




