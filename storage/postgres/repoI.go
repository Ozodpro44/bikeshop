package postgres

import (
	"bikeshop/models"
	"context"
)

type ProductRepoI interface {
	CreateProduct(ctx context.Context, product *models.Bike) (*models.Bike, error)
	GetProducts(ctx context.Context, page, limit int32) (*models.GetBikesList, error)
	GetProduct(ctx context.Context, id string) (*models.Bike, error)
	UpdateProduct(ctx context.Context, product *models.Bike) (*models.Bike, error)
	DeleteProduct(ctx context.Context, id string) error
}