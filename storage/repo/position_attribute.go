package repo

import "profession_service/genproto/profession_service"

type PositionAttributeRepoI interface {
	Create(req *profession_service.CreatePositionAttribute) (string, error)
	GetAll(req *profession_service.GetAllPositionAttributeRequest) (*profession_service.GetAllPositionAttributeResponse, error)
	GetById(id string) (*profession_service.GetPositionAttribute, error)
	Update(req *profession_service.PositionAttribute) (*profession_service.AfterPositionAttributeUpdate, error)
	Delete(id string) (*profession_service.IsDeletedPA, error)
}
