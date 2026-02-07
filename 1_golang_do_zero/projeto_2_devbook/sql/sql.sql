CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
    
)ENGINE=InnoDB;

CREATE TABLE seguidores (
    seguido_id INT NOT NULL,  -- o usuário que está sendo seguido (alvo)
    seguidor_id INT NOT NULL, -- o usuário que está seguindo (quem faz a ação)
    PRIMARY KEY (seguido_id, seguidor_id),
    FOREIGN KEY (seguido_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE
) ENGINE=InnoDB;