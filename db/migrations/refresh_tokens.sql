CREATE TABLE IF NOT EXISTS rep_user_refresh_tokens
(
    id             BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of refresh token.',
    token          VARCHAR(64)     NOT NULL COMMENT 'Token.',
    user_id        BIGINT UNSIGNED NOT NULL COMMENT 'Id of user.',
    expires_at     TIMESTAMP       NOT NULL COMMENT 'Status of expiration refresh token.',
    created_at     TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation refresh token',
    updated_at     TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating refresh token',

    PRIMARY KEY (id),
    CONSTRAINT rep_refresh_tokens_token_uk UNIQUE (token)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Refresh tokens';