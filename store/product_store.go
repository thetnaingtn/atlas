package store

import (
	"context"
	"time"
)

// CreateProduct stores a new product in the database after applying business logic.
func (s *Store) CreateProduct(ctx context.Context, p *Product) (*Product, error) {
	if p == nil {
		return nil, nil
	}
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	return s.driver.CreateProduct(ctx, p)
}

// UpdateProduct updates an existing product in the database.
func (s *Store) UpdateProduct(ctx context.Context, p *Product) (*Product, error) {
	if p == nil {
		return nil, nil
	}
	p.UpdatedAt = time.Now()
	return s.driver.UpdateProduct(ctx, p)
}

// ListProducts retrieves all products from the database.
func (s *Store) ListProducts(ctx context.Context) ([]*Product, error) {
	return s.driver.ListProducts(ctx)
}

// DeleteProduct removes a product by ID.
func (s *Store) DeleteProduct(ctx context.Context, id int64) error {
	return s.driver.DeleteProduct(ctx, id)
}
