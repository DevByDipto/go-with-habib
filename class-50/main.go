package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func aboutHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world, i am about")
}

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader((200))
	}
	if r.Method != "GET" {
		http.Error(w, "plz give me GET request", 400)
	}


	sendData(w,productList,200)
}

// createProductHandler
func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)
	if r.Method != "POST" {
		http.Error(w, "Plz give me POST request", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	
	sendData(w,newProduct,201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET,POST,PUT,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Habib")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader((200))
	}
}

func sendData(w http.ResponseWriter, data interface{},statusCode int){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandeler)
	mux.HandleFunc("/about", aboutHandeler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/creat-products", createProduct)

	fmt.Println("server runing 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init(){
	// product1:=
}

/*
1. handeler kii ?
ans: kono akta path mathch howar por oi math er jonno j actual function ke handle kora hoi ta handler
2. handler /controller akoi kotha
*/