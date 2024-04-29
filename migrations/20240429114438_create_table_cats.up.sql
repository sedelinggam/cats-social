CREATE TYPE cat_sex AS ENUM ('male', 'female');

CREATE TYPE cat_race AS ENUM ('Persian','Maine Coon','Siamese','Ragdoll','Bengal','Sphynx','British Shorthair','Abyssinian','Scottish Fold','Birman');

CREATE TABLE cats (
  id uuid PRIMARY KEY,
  user_id uuid NOT NULL,
  name varchar(30) NOT NULL,
  race cat_race NOT NULL,
  sex cat_sex NOT NULL,
  age_in_month smallint NOT NULL,
  image_urls text[] NOT NULL,
  description varchar(200) NOT NULL,
  is_already_matched boolean NOT NULL,
  created_at timestamptz,
  deleted_at timestamptz,
  update_at timestamptz,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);