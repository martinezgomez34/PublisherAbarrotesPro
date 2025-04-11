package repositories

import (
	"publisher/src/product/domain"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) SaveProduct(product domain.Product) error {
	query := `INSERT INTO products (name, description, price) VALUES (?, ?, ?)`
	_, err := r.DB.Exec(query, product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) EditProduct(product domain.Product) error {
	query := `UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?`
	_, err := r.DB.Exec(query, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(productID string) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := r.DB.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) GetAll() ([]*domain.Product, error) {
	query := `SELECT id, name, description, price FROM products`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	
	return products, nil
}

func (r *ProductRepository) GetByID(id string) (*domain.Product, error) {
	query := `SELECT id, name, description, price FROM products WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	var product domain.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) ProductExists(productID string) bool {
	query := `SELECT COUNT(1) FROM products WHERE id = ?`
	var count int
	err := r.DB.QueryRow(query, productID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
