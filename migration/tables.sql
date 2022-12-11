DROP TABLE IF EXISTS staff CASCADE;
CREATE TABLE staff (
                       id bigserial PRIMARY KEY,
                       fname VARCHAR(15) NOT NULL,
                       lname VARCHAR(15) NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password_hash bytea NOT NULL,
                       role varchar(6) );

DROP TABLE IF EXISTS questions CASCADE;
CREATE TABLE questions (
                       id bigserial PRIMARY KEY,
                       created_at timestamp(0) with time zone NOT NULL,
                       question text NOT NULL,
                       status bool,
                       staff_id bigint NOT NULL REFERENCES staff ON DELETE CASCADE   )         ;


DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
                          token TEXT PRIMARY KEY,
                          data BYTEA NOT NULL,
                          expiry TIMESTAMPTZ NOT NULL ) ;

CREATE INDEX sessions_expiry_idx ON sessions (expiry) ;