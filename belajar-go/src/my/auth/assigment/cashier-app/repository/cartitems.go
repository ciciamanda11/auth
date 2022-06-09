package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	var sqlStatement string
	var cartItems []CartItem

	sqlStatement = `SELECT
						c.id,
						c.product_id,
						c.quantity,
						p.product_name,
						p.price
					FROM cart_items c
					INNER JOIN products p ON c.product_id = p.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem CartItem
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.ProductID,
			&cartItem.Quantity,
			&cartItem.ProductName,
			&cartItem.Price)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	var sqlStatement string

	sqlStatement = `SELECT
						c.id,
						c.product_id,
						c.quantity,
						p.product_name,
						p.price
					FROM cart_items c
					INNER JOIN products p ON c.product_id = p.id
					WHERE c.product_id = ?`

	row := c.db.QueryRow(sqlStatement, productID)
	err := row.Scan(
		&cartItem.ID,
		&cartItem.ProductID,
		&cartItem.Quantity,
		&cartItem.ProductName,
		&cartItem.Price)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {

	sqlStatement := `INSERT INTO cart_items (product_id, quantity) VALUES (?, ?)`

	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {

	sqlStatement := `UPDATE cart_items SET quantity = quantity + ? WHERE product_id = ?`

	_, err := c.db.Exec(sqlStatement, cartItem.Quantity, cartItem.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) ResetCartItems() error {

	sqlStatement := `DELETE FROM cart_items`

	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string

	sqlStatement = `SELECT
						SUM(p.price * c.quantity)
					FROM cart_items c
					INNER JOIN products p ON c.product_id = p.id`

	var totalPrice int
	row := c.db.QueryRow(sqlStatement)
	err := row.Scan(&totalPrice)
	if err != nil {
		return 0, err
	}

	return totalPrice, nil
}
