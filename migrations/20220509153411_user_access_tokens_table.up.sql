CREATE TABLE user_access_tokens (
    id INT NOT NULL AUTO_INCREMENT ,
    token TEXT  NULL ,
    exp_date DATETIME NULL ,
    created_at DATETIME NULL ,
    PRIMARY KEY (id)
) ENGINE = InnoDB CHAR SET utf8;

CREATE INDEX token_idx1 ON user_access_tokens(token(255)) USING BTREE;