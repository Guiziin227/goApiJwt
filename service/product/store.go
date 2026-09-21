package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/guiziin227/goApiJwt/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]*types.Product, error) {

	rows, err := s.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		prod, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}

	return products, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	row := s.db.QueryRow("SELECT * FROM products WHERE id = ?", id)

	prod := new(types.Product)
	if err := row.Scan(
		&prod.ID,
		&prod.Name,
		&prod.Description,
		&prod.Image,
		&prod.Price,
		&prod.Quantity,
		&prod.CreatedAt,
	); err != nil {
		return nil, err
	}

	return prod, nil
}

func (s *Store) GetProductsByID(ids []int) ([]types.Product, error) {
	if len(ids) == 0 {
		return []types.Product{}, nil
	}

	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := s.db.Query(
		fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", placeholders),
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]types.Product, 0, len(ids))
	for rows.Next() {
		prod, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *prod)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec(
		"UPDATE products SET name = ?, description = ?, image = ?, price = ?, quantity = ? WHERE id = ?",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
		product.ID,
	)
	return err
}

func (s *Store) CreateProduct(product *types.Product) error {
	_, err := s.db.Exec("INSERT INTO products "+
		"(name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity)

	if err != nil {
		return err
	}
	return nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	prod := new(types.Product)

	err := rows.Scan(
		&prod.ID,
		&prod.Name,
		&prod.Description,
		&prod.Image,
		&prod.Price,
		&prod.Quantity,
		&prod.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return prod, nil
}
