CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT ,
    first_name VARCHAR(200) NULL ,
    last_name VARCHAR(200) NULL ,
    email TEXT NULL ,
    password TEXT NULL ,
    access_level BOOLEAN NULL DEFAULT  1,
    created_at DATETIME NULL,
    updated_at DATETIME NULL ,
    PRIMARY KEY (id)
) ENGINE = InnoDB;