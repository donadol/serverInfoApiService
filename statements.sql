CREATE USER IF NOT EXISTS maxroach;

CREATE DATABASE IF NOT EXISTS serversInfo;

GRANT ALL ON DATABASE serversInfo TO maxroach;

USE serversinfo;

CREATE TABLE IF NOT EXISTS DOMAIN (
  id SERIAL,
  consulted_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  host TEXT NOT NULL,
  PRIMARY KEY (id));

CREATE TABLE IF NOT EXISTS INFOSERVER (
  id SERIAL,
  servers_changed BOOLEAN NOT NULL,
  ssl_grade VARCHAR(3) NOT NULL,
  previous_ssl_grade VARCHAR(3) NOT NULL,
  logo TEXT NOT NULL,
  title TEXT NOT NULL,
  is_down BOOLEAN NOT NULL,
  domain_id INT NOT NULL,
  PRIMARY KEY (id),
  INDEX domain_id_idx (domain_id ASC),
  CONSTRAINT domain_id
    FOREIGN KEY (domain_id)
    REFERENCES domain (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

CREATE TABLE IF NOT EXISTS SERVER (
  id SERIAL,
  address TEXT NOT NULL,
  ssl_grade VARCHAR(3) NOT NULL,
  country VARCHAR(5) NOT NULL,
  owner TEXT NOT NULL,
  infoserver_id INT NOT NULL,
  PRIMARY KEY (id),
  INDEX infoserver_id_idx (infoserver_id ASC),
  CONSTRAINT infoserver_id
    FOREIGN KEY (infoserver_id)
    REFERENCES infoServer (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

\q