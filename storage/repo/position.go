package repo

import "profession_service/genproto/profession_service"

type PositionRepoI interface {
	Create(req *profession_service.CreatePosition) (string, error)
	GetAll(req *profession_service.GetAllPositionRequest) (*profession_service.GetAllPositionResponse, error)
	GetById(id string) (*profession_service.GetPosition, error)
	Update(req *profession_service.Position) (*profession_service.AfterUpdatePosition, error)
	Delete(id string) (*profession_service.IsDeleted, error)
}
