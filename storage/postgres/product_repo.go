package postgres

import (
	"bikeshop/models"
	log "bikeshop/pkg/logger"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
	log log.Log
}

func NewProductRepo(db *pgxpool.Pool, log log.Log) ProductRepoI {
	return &productRepo{
		db, log,
	}
}

func (p *productRepo) CreateProduct(ctx context.Context, product *models.Bike) (*models.Bike, error) {
	p.log.Debug("Request in CreateProduct..")

	
	return nil ,nil
}
func (p *productRepo) GetProducts(ctx context.Context, page, limit int32) (*models.GetBikesList, error) {
	return nil, nil
}
func (p *productRepo) GetProduct(ctx context.Context, id string) (*models.Bike, error) {
	return nil, nil
}
func (p *productRepo) UpdateProduct(ctx context.Context, product *models.Bike) (*models.Bike, error) {
	return nil, nil
}
func (p *productRepo) DeleteProduct(ctx context.Context, id string) error {
	return nil
}	