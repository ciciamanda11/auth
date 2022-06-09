package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {

	var product Product
	err := p.db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.ID, &product.Category, &product.ProductName, &product.Price, &product.Quantity)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {

	var product Product
	err := p.db.QueryRow("SELECT * FROM products WHERE product_name = ?", productName).Scan(
		&product.ID,
		&product.Category,
		&product.ProductName,
		&product.Price,
		&product.Quantity)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {

	var products []Product
	rows, err := p.db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Category,
			&product.Price,
			&product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
