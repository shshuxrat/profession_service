package storage

import (
	"profession_service/storage/postgres"
	"profession_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Position() repo.PositionRepoI
	PositionAttribute() repo.PositionAttributeRepoI
}

type storagePG struct {
	position          repo.PositionRepoI
	positionAttribute repo.PositionAttributeRepoI
}

func NewStroragePG(db *sqlx.DB) StorageI {
	return &storagePG{
		position:          postgres.NewPositionRepo(db),
		positionAttribute: postgres.NewPositionAttributeRepo(db),
	}
}

func (s *storagePG) Position() repo.PositionRepoI {
	return s.position
}

func (s *storagePG) PositionAttribute() repo.PositionAttributeRepoI {
	return s.positionAttribute
}
