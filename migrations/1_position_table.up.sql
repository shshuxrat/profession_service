CREATE TABLE IF NOT EXISTS position(
    id uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    profession_id uuid REFERENCES profession(id) ON DELETE CASCADE,
    company_id uuid  REFERENCES company(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT(Now()),
	updated_at TIMESTAMP DEFAULT(Now())
);
