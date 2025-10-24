CREATE TABLE IF NOT EXISTS rep_transaction_types
(
    id         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of transaction type.',
    name       VARCHAR(64) NOT NULL COMMENT 'Name of transaction type.',
    direction  TINYINT(1)  NOT NULL DEFAULT 0 COMMENT 'Direction of transaction type (1 - in, 0 - out)',
    regularity TINYINT(1)  NOT NULL DEFAULT 0 COMMENT 'Regularity of transaction type (1 - regular, 0 - one-time)',
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation transaction type',
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating transaction type',

    PRIMARY KEY (id),
    CONSTRAINT rep_transaction_types_name_direction_regularity_uk UNIQUE (name, direction, regularity)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Transaction types';