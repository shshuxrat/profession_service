package postgres

import (
	"profession_service/genproto/profession_service"
	"profession_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type positionRepo struct {
	db *sqlx.DB
}

func NewPositionRepo(db *sqlx.DB) repo.PositionRepoI {
	return &positionRepo{
		db: db,
	}
}

func (p *positionRepo) Create(req *profession_service.CreatePosition) (string, error) {
	var (
		id uuid.UUID
	)
	tx, err := p.db.Begin()
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

	query := `INSERT INTO position(id, name, profession_id,company_id) VALUES($1,$2,$3,$4)`

	_, err = p.db.Exec(query, id, req.Name, req.ProfessionId, req.CompanyId)

	if err != nil {
		return "", nil
	}

	return id.String(), nil
}

func (p *positionRepo) GetAll(req *profession_service.GetAllPositionRequest) (*profession_service.GetAllPositionResponse, error) {
	var (
		filter    string
		count     int32
		positions []*profession_service.Position
	)

	args := make(map[string]interface{})

	if req.Name != "" {
		filter += ` AND name ILIKE '%' || :filter_name || '%'`
		args["filter_name"] = req.Name
	}

	filter += ` LIMIT :lim OFFSET :off `
	args["lim"] = req.Limit
	args["off"] = req.Offset
	countQuery := `SELECT count(1) FROM position WHERE true` + filter
	rows, err := p.db.NamedQuery(countQuery, args)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	query := `SELECT id, name, profession_id, company_id,created_at,updated_at FROM position WHERE true ` + filter

	rows, err = p.db.NamedQuery(query, args)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var position profession_service.Position
		err = rows.Scan(&position.Id, &position.Name, &position.ProfessionId, &position.CompanyId,
			&position.CreatedAt, &position.UpdatedAt)
		if err != nil {
			return nil, err
		}
		positions = append(positions, &position)
	}

	return &profession_service.GetAllPositionResponse{
		Positions: positions,
		Count:     count,
	}, nil
}

func (p *positionRepo) GetById(id string) (*profession_service.GetPosition, error) {
	query := `SELECT ps.id, ps.name, ps.profession_id, ps.company_id,ps.created_at, ps.updated_at, pr.name, c.name
			FROM position AS ps, profession AS pr,  company AS c
			WHERE ps.company_id = c.id AND ps.profession_id = pr.id  AND ps.id = $1`
	rows, err := p.db.Query(query, id)

	if err != nil {
		return nil, err
	}
	var companyName, professionName string
	var position profession_service.Position
	for rows.Next() {

		err = rows.Scan(&position.Id, &position.Name, &position.ProfessionId, &position.CompanyId,
			&position.CreatedAt, &position.UpdatedAt, &professionName, &companyName)
		if err != nil {
			return nil, err
		}
	}

	return &profession_service.GetPosition{
		Position:   &position,
		Company:    companyName,
		Profession: professionName,
	}, nil
}

func (p *positionRepo) Update(req *profession_service.Position) (*profession_service.AfterUpdatePosition, error) {
	old, err := p.GetById(req.Id)

	if err != nil {
		return nil, err
	}

	_, err = p.db.Exec(
		`UPDATE position 
		SET name=$1 , profession_id=$2, company_id=$3, updated_at =Now()
		WHERE id=$4`,
		req.Name,
		req.ProfessionId,
		req.CompanyId,
		req.Id,
	)

	if err != nil {
		return nil, err
	}

	newone, err := p.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return &profession_service.AfterUpdatePosition{
		Old: old,
		New: newone,
	}, nil
}

func (p *positionRepo) Delete(id string) (*profession_service.IsDeleted, error) {
	_, err := p.db.Exec(
		`DELETE FROM position WHERE id = $1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &profession_service.IsDeleted{
		IsDeleted: "YES",
	}, nil
}
