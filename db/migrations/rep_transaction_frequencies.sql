CREATE TABLE IF NOT EXISTS rep_transaction_frequencies
(
    id         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of transaction frequency.',
    name       VARCHAR(64) NOT NULL COMMENT 'Name of transaction frequency.',
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation transaction frequency',
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating transaction frequency',

    PRIMARY KEY (id),
    CONSTRAINT rep_transaction_frequencies_name_uk UNIQUE (name)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Transaction frequencies';