package database

import (
	"database/sql"

	"github.com/marcosgomesdev/goapi/internal/entity"
)

type Product struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *Product {
	return &Product{db: db}
}

func (pd *Product) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, price, category_id, image_url, description FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL, &product.Description); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}
	return products, nil
}

func (pd *Product) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *Product) GetProductByCategory(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, price, category_id, image_url FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}
	return products, nil
}

func (pd *Product) CreateProduct(product *entity.Product) (string, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)

	if err != nil {
		return "", err
	}

	return product.ID, nil

}
