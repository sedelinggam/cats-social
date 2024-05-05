CREATE TABLE matches(
  id uuid PRIMARY KEY,
  match_cat_id uuid NOT NULL,
  user_cat_id uuid NOT NULL,
  created_at timestamptz,
  message varchar(120) NOT NULL,

  CONSTRAINT fk_match_cat_id FOREIGN KEY(match_cat_id) REFERENCES cats(id),
  CONSTRAINT fk_user_cat_id FOREIGN KEY(user_cat_id) REFERENCES cats(id)
);

CREATE INDEX matches_id ON matches (id);
CREATE INDEX matches_match_cat_id ON matches (match_cat_id);
CREATE INDEX matches_match_user_cat_id ON matches (user_cat_id);