ALTER TABLE matches
ADD COLUMN issuer_id uuid REFERENCES users(id),
ALTER COLUMN issuer_id SET NOT NULL;