CREATE TABLE IF NOT EXISTS rep_users
(
    id                 BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of user.',
    email              VARCHAR(64)  NOT NULL COMMENT 'Email address of user.',
    password           VARCHAR(255) NOT NULL COMMENT 'Password.',

    name               VARCHAR(64)           DEFAULT NULL COMMENT 'Name/Nickname',
    phone              VARCHAR(16)           DEFAULT NULL COMMENT 'Phone number',

    email_confirm      TINYINT(1)   NOT NULL DEFAULT 0 COMMENT 'Status of confirmation email.',
    email_confirm_code VARCHAR(64)           DEFAULT NULL COMMENT 'Code for confirmation email.',
    email_confirmed_at TIMESTAMP             DEFAULT NULL COMMENT 'Date of confirmation email',

    deleted            TINYINT(1)   NOT NULL DEFAULT 0 COMMENT 'Status of deletion user.',
    created_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation user',
    updated_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating user',

    PRIMARY KEY (id),
    CONSTRAINT rep_users_email_uk UNIQUE (email)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Users';