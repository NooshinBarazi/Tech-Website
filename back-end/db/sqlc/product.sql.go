// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: product.sql

package db

import (
	"context"
)

const getProduct = `-- name: GetProduct :one
SELECT id, product_image, product_name, product_category, sold_number, brand, color, screen_size, hard_disk_size, display, graphic, processor, in_the_box, height, width, cost, star FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProductImage,
		&i.ProductName,
		&i.ProductCategory,
		&i.SoldNumber,
		&i.Brand,
		&i.Color,
		&i.ScreenSize,
		&i.HardDiskSize,
		&i.Display,
		&i.Graphic,
		&i.Processor,
		&i.InTheBox,
		&i.Height,
		&i.Width,
		&i.Cost,
		&i.Star,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, product_image, product_name, product_category, sold_number, brand, color, screen_size, hard_disk_size, display, graphic, processor, in_the_box, height, width, cost, star FROM products
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.ProductImage,
			&i.ProductName,
			&i.ProductCategory,
			&i.SoldNumber,
			&i.Brand,
			&i.Color,
			&i.ScreenSize,
			&i.HardDiskSize,
			&i.Display,
			&i.Graphic,
			&i.Processor,
			&i.InTheBox,
			&i.Height,
			&i.Width,
			&i.Cost,
			&i.Star,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
  set star = $2
WHERE id = $1
RETURNING id, product_image, product_name, product_category, sold_number, brand, color, screen_size, hard_disk_size, display, graphic, processor, in_the_box, height, width, cost, star
`

type UpdateProductParams struct {
	ID   int64   `json:"id"`
	Star float64 `json:"star"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct, arg.ID, arg.Star)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProductImage,
		&i.ProductName,
		&i.ProductCategory,
		&i.SoldNumber,
		&i.Brand,
		&i.Color,
		&i.ScreenSize,
		&i.HardDiskSize,
		&i.Display,
		&i.Graphic,
		&i.Processor,
		&i.InTheBox,
		&i.Height,
		&i.Width,
		&i.Cost,
		&i.Star,
	)
	return i, err
}
