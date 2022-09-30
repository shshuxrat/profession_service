package service

import (
	"context"
	"profession_service/genproto/profession_service"
	"profession_service/pkg/logger"
	"profession_service/storage"

	"github.com/jmoiron/sqlx"
)

type positionAttributeService struct {
	logger  logger.LoggerI
	storage storage.StorageI
	profession_service.UnimplementedPositionAttributeServiceServer
}

func NewPositionAttributeService(log logger.LoggerI, db *sqlx.DB) *positionAttributeService {
	return &positionAttributeService{
		logger:  log,
		storage: storage.NewStroragePG(db),
	}
}

func (p *positionAttributeService) Create(c context.Context, req *profession_service.CreatePositionAttribute) (*profession_service.PositionAttributeId, error) {
	id, err := p.storage.PositionAttribute().Create(req)

	if err != nil {
		return nil, err
	}

	return &profession_service.PositionAttributeId{
		Id: id,
	}, nil
}

func (p *positionAttributeService) GetAll(c context.Context, req *profession_service.GetAllPositionAttributeRequest) (*profession_service.GetAllPositionAttributeResponse, error) {
	resp, err := p.storage.PositionAttribute().GetAll(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *positionAttributeService) GetById(c context.Context, req *profession_service.PositionAttributeId) (*profession_service.GetPositionAttribute, error) {
	positionAttribute, err := p.storage.PositionAttribute().GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return positionAttribute, nil
}

func (p *positionAttributeService) Update(c context.Context, req *profession_service.PositionAttribute) (*profession_service.AfterPositionAttributeUpdate, error) {
	resp, err := p.storage.PositionAttribute().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (p *positionAttributeService) Delete(c context.Context, req *profession_service.PositionAttributeId) (*profession_service.IsDeletedPA, error) {
	resp, err := p.storage.PositionAttribute().Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
