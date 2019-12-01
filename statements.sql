CREATE USER IF NOT EXISTS maxroach;

CREATE DATABASE serversInfo;

GRANT ALL ON DATABASE serversInfo TO maxroach;

USE serversinfo;

CREATE TABLE IF NOT EXISTS DOMAIN (
  id INT NOT NULL AUTO_INCREMENT,
  last_consult DATETIME NOT NULL,
  host VARCHAR(45) COLLATE 'DEFAULT' NOT NULL,
  PRIMARY KEY (id));

CREATE TABLE IF NOT EXISTS INFOSERVER (
  id INT NOT NULL AUTO_INCREMENT,
  servers_changed BINARY(1) NOT NULL,
  ssl_grade VARCHAR(3) NOT NULL,
  previous_ssl_grade VARCHAR(3) NOT NULL,
  logo VARCHAR(45) NOT NULL,
  title VARCHAR(45) NOT NULL,
  is_down BINARY(1) NOT NULL,
  domain_id INT NOT NULL,
  PRIMARY KEY (id),
  INDEX domain_id_idx (domain_id ASC) VISIBLE,
  CONSTRAINT domain_id
    FOREIGN KEY (domain_id)
    REFERENCES domain (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

CREATE TABLE IF NOT EXISTS SERVER (
  id INT NOT NULL AUTO_INCREMENT,
  address VARCHAR(45) NOT NULL,
  ssl_grade VARCHAR(3) NOT NULL,
  country VARCHAR(5) NOT NULL,
  owner VARCHAR(45) NOT NULL,
  infoserver_id INT NOT NULL,
  PRIMARY KEY (id),
  INDEX infoserver_id_idx (infoserver_id ASC) VISIBLE,
  CONSTRAINT infoserver_id
    FOREIGN KEY (infoserver_id)
    REFERENCES infoServer (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

\q