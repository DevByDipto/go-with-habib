package repo

import (
	"errors"
)

// Product represents the data model for a store item
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

// ProductRepo defines the behavior of our product storage
type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// NewProductRepo is a constructor that initializes the repo with seed data
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

// --- CRUD Methods ---

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product
	found := false

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("cannot delete: product not found")
	}

	r.productList = tempList
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
			return &product, nil
		}
	}
	return nil, errors.New("cannot update: product not found")
}

// --- Seed Data ---

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. I love orange.",
		Price:       100,
		ImgUrl:      "https://www.dole.com/sites/default/files/media/2025-0",
	}

	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green. I hate apple.",
		Price:       40,
		ImgUrl:      "https://www.harrisfarm.com.au/cdn/shop/products/40715",
	}

	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is boring. I feel bored eating banana.",
		Price:       5,
		ImgUrl:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAW",
	}

	prd4 := &Product{
		ID:          4, // Fixed ID from screenshot (was 3)
		Title:       "Banana Clone",
		Description: "Another boring banana.",
		Price:       5,
		ImgUrl:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAW",
	}

	prd5 := &Product{
		ID:          5, // Fixed ID from screenshot (was 3)
		Title:       "Angur Fol",
		Description: "Grapes are great.",
		Price:       140,
		ImgUrl:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAW",
	}

	prd6 := &Product{
		ID:          6, // Fixed ID from screenshot (was 3)
		Title:       "Mango",
		Description: "Mango is my favorite. I love it very much.",
		Price:       10000000,
		ImgUrl:      "https://www.dole.com/sites/default/files/styles/512w3",
	}

	r.productList = append(r.productList, prd1, prd2, prd3, prd4, prd5, prd6)
}