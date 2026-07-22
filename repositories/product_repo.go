package repositories

import (
	"database/sql"
	"golang_standart_project/models"
)

type ProductRepository interface {
	Insert(product models.Product) (int, error)
	FindAll() ([]models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Insert(p models.Product) (int, error) {
	query := "INSERT INTO products (product_name, description, price, stock_quantity) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, p.ProductName, p.Description, p.Price, p.StockQuantity)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	query := "SELECT id, product_name, description, price, stock_quantity FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.ProductName, &p.Description, &p.Price, &p.StockQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
