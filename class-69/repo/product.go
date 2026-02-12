package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)


type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// NewProductRepo is a constructor that initializes the repo with seed data
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}


// --- CRUD Methods ---

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO domain.s (
			title,
			description,
			price,
			img_url
		) VALUES (
			$1, $2, $3, $4
		)
		RETURNING id
	`

	// Execute the query and scan the returned ID back into the product struct
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product

	query := `
		SELECT 
			id, 
			title, 
			description, 
			price, 
			img_url 
		FROM products 
		WHERE id = $1
	`

	// r.db.Get সরাসরি কুয়েরি রেজাল্টকে prd ভ্যারিয়েবলে ম্যাপ করে ফেলে
	err := r.db.Get(&prd, query, id)
	if err != nil {
		// যদি ডাটাবেসে এই আইডি দিয়ে কোন রো না পাওয়া যায়
		if err == sql.ErrNoRows {
			return nil, nil
		}
		// অন্য কোন এরর হলে সেটা রিটার্ন করবে
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	var prdList []*domain.Product // slice -> address, cap, len

	query := `
		SELECT 
			id, 
			title, 
			description, 
			price, 
			img_url 
		FROM products
	`

	// r.db.Select ব্যবহার করা হয়েছে কারণ আমরা একাধিক রো (Rows) আশা করছি
	err := r.db.Select(&prdList, query)
	if err != nil {
		return nil, err
	}

	return prdList, nil
}

func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`

	// Exec ব্যবহার করা হয় যখন আমরা কোন ডাটা রিটার্ন চাই না (যেমন: Delete, Update)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET title=$1, description=$2, price=$3, img_url=$4
		WHERE id = $5
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &p, nil
}


