package postgres

import (
	"fmt"
	"profession_service/genproto/profession_service"
	"profession_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type positionAttributeRepo struct {
	db *sqlx.DB
}

func NewPositionAttributeRepo(db *sqlx.DB) repo.PositionAttributeRepoI {
	return &positionAttributeRepo{
		db: db,
	}

}

func (pa *positionAttributeRepo) Create(req *profession_service.CreatePositionAttribute) (string, error) {
	var (
		id uuid.UUID
	)
	tx, err := pa.db.Begin()
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	id, err = uuid.NewRandom()

	if err != nil {
		return "", nil
	}

	query := `INSERT INTO position_attribute(id,value, attribute_id,position_id) VALUES($1,$2,$3,$4)`

	_, err = pa.db.Exec(query, id, req.Value, req.AttributeId, req.PositionId)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return id.String(), nil
}

func (pa *positionAttributeRepo) GetAll(req *profession_service.GetAllPositionAttributeRequest) (*profession_service.GetAllPositionAttributeResponse, error) {
	var (
		filter string
		count  int32
		arr    []*profession_service.PositionAttribute
	)

	args := make(map[string]interface{})

	if req.Value != "" {
		filter += ` AND value ILIKE '%' || :filter_value || '%' `
		args["filter_value"] = req.Value
	}
	filter += ` LIMIT :lim OFFSET :off `
	args["lim"] = req.Limit
	args["off"] = req.Offset

	countQuery := `SELECT count(1) FROM position_attribute WHERE true` + filter
	rows, err := pa.db.NamedQuery(countQuery, args)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	query := `SELECT id, value , attribute_id, position_id, created_at, updated_at
			FROM position_attribute 
			WHERE true` + filter
	rows, err = pa.db.NamedQuery(query, args)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pos_att profession_service.PositionAttribute
		err = rows.Scan(&pos_att.Id, &pos_att.Value, &pos_att.AttributeId, &pos_att.PositionId,
			&pos_att.CreatedAt, &pos_att.UpdatedAt)
		if err != nil {
			return nil, err
		}

		arr = append(arr, &pos_att)
	}

	return &profession_service.GetAllPositionAttributeResponse{
		PositionAttributes: arr,
		Count:              count,
	}, nil
}

func (pa *positionAttributeRepo) GetById(id string) (*profession_service.GetPositionAttribute, error) {

	query := `SELECT pa.id, pa.value, pa.position_id, pa.attribute_id, pa.created_at, pa.updated_at, a.name, p.name
			FROM position_attribute AS pa, attribute AS a, position AS p
			WHERE pa.position_id = p.id AND pa.attribute_id = a.id  AND pa.id = $1`
	row, err := pa.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var (
		pp            profession_service.PositionAttribute
		attributeName string
		positionName  string
	)
	for row.Next() {
		err = row.Scan(&pp.Id, &pp.Value, &pp.PositionId, &pp.AttributeId,
			&pp.CreatedAt, &pp.UpdatedAt, &attributeName, &positionName)
		if err != nil {
			return nil, err
		}
	}

	return &profession_service.GetPositionAttribute{
		PositionAttribute: &pp,
		Attribute:         attributeName,
		Position:          positionName,
	}, nil

}

func (pa *positionAttributeRepo) Update(req *profession_service.PositionAttribute) (*profession_service.AfterPositionAttributeUpdate, error) {
	old, err := pa.GetById(req.Id)
	if err != nil {
		return nil, err
	}
	_, err = pa.db.Exec(
		`UPDATE position_attribute
		SET value=$1 , attribute_id=$2, position_id=$3, updated_at = Now()
		WHERE id=$4`,
		req.Value,
		req.AttributeId,
		req.PositionId,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	newOne, err := pa.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return &profession_service.AfterPositionAttributeUpdate{
		Old: old,
		New: newOne,
	}, nil

}

func (pa *positionAttributeRepo) Delete(id string) (*profession_service.IsDeletedPA, error) {
	_, err := pa.db.Exec(
		`DELETE FROM position_attribute WHERE id = $1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &profession_service.IsDeletedPA{
		IsDeleted: "YES",
	}, nil
}
