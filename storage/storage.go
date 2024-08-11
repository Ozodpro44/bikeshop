package storage

import (
	log "bikeshop/pkg/logger"
	"bikeshop/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageI interface {
	GetProductRepo() postgres.ProductRepoI
}

type storage struct {
	productRepo postgres.ProductRepoI
}

func NewStorage(db *pgxpool.Pool, log log.Log) StorageI {
	return &storage{
		productRepo: postgres.NewProductRepo(db, log),
	}
}

func (s *storage) GetProductRepo() postgres.ProductRepoI {
	return s.productRepo
}