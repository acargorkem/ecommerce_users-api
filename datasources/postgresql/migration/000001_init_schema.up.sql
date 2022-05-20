CREATE TABLE IF NOT EXISTS "users" (
  id bigserial NOT NULL,
  first_name character varying(100) NOT NULL,
  last_name character varying(100) NOT NULL,
  email character varying(100) UNIQUE NOT NULL,
  status character varying(100) NOT NULL,
  hashed_password character varying(255) NOT NULL,
  created_at character varying(100) NOT NULL,
  PRIMARY KEY (id)
);
