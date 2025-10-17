CREATE TABLE IF NOT EXISTS rep_estate_types
(
    id         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of estate type.',
    name       VARCHAR(64) NOT NULL COMMENT 'Name of estate type.',
    active     TINYINT(1)  NOT NULL DEFAULT 0 COMMENT 'Status of deletion estate type.',
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation estate type',
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating estate type',

    PRIMARY KEY (id),
    CONSTRAINT rep_estate_types_name_uk UNIQUE (name)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Estate types';