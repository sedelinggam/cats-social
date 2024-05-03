ALTER TABLE matches
ADD COLUMN receiver_id uuid REFERENCES users(id),
ALTER COLUMN receiver_id SET NOT NULL;