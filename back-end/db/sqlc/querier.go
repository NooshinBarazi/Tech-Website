// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, categoryName string) (Category, error)
	CreateImage(ctx context.Context, arg CreateImageParams) (Image, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
}

var _ Querier = (*Queries)(nil)
