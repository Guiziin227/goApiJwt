package product

import (
	"database/sql"

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
