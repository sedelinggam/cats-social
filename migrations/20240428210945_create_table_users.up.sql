CREATE TABLE users (
  user_id uuid PRIMARY KEY, 
  email VARCHAR  UNIQUE NOT NULL, 
  password VARCHAR (72) NOT NULL,
  name VARCHAR (50) NOT NULL, 
  created_at timestamptz NOT NULL
);