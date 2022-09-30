CREATE TABLE IF NOT EXISTS position_attribute(
    id uuid PRIMARY KEY,
    "value" VARCHAR(255) NOT NULL,
    attribute_id uuid REFERENCES attribute(id) ON DELETE CASCADE,
    position_id uuid  REFERENCES position(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT(Now()),
	updated_at TIMESTAMP DEFAULT(Now()),
    deleted_at TIMESTAMP
);
