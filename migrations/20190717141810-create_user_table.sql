
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id INTEGER AUTO_INCREMENT NOT NULL,
    login VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    recovery_token text,
    PRIMARY KEY(id)
);

-- password hash = "password" 
INSERT INTO users(login, email, password) VALUES('admin', 'kpi_lithium@ukr.net', '$2a$10$m4O6aon19..QGP8rX9vriunnjseBaIjVrXudHE37lR5aH.lrIgYIi');

-- +migrate Down
DROP TABLE users;