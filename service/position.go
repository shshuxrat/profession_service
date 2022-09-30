package service

import (
	"context"
	"profession_service/genproto/profession_service"
	"profession_service/pkg/logger"
	"profession_service/storage"

	"github.com/jmoiron/sqlx"
)

type positionService struct {
	logger  logger.LoggerI
	storage storage.StorageI
	profession_service.UnimplementedPositionServiceServer
}

func NewPositionService(log logger.LoggerI, db *sqlx.DB) *positionService {
	return &positionService{
		logger:  log,
		storage: storage.NewStroragePG(db),
	}
}

func (p *positionService) Create(c context.Context, req *profession_service.CreatePosition) (*profession_service.PositionId, error) {
	positionId, err := p.storage.Position().Create(req)
	if err != nil {
		return nil, err
	}
	return &profession_service.PositionId{
		Id: positionId,
	}, nil
}

func (p *positionService) GetAll(c context.Context, req *profession_service.GetAllPositionRequest) (*profession_service.GetAllPositionResponse, error) {
	positions, err := p.storage.Position().GetAll(req)
	if err != nil {
		return nil, err
	}

	return positions, nil

}

func (p *positionService) GetById(C context.Context, req *profession_service.PositionId) (*profession_service.GetPosition, error) {
	position, err := p.storage.Position().GetById(req.Id)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (p *positionService) Update(c context.Context, req *profession_service.Position) (*profession_service.AfterUpdatePosition, error) {
	positions, err := p.storage.Position().Update(req)
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (p *positionService) Delete(c context.Context, req *profession_service.PositionId) (*profession_service.IsDeleted, error) {
	isdeleted, err := p.storage.Position().Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return isdeleted, nil

}
