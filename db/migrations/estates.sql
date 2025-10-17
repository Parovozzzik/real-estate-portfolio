CREATE TABLE IF NOT EXISTS rep_estates
(
    id             BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of estate.',
    estate_type_id BIGINT UNSIGNED NOT NULL COMMENT 'Id of estate type.',
    name           VARCHAR(64)     NOT NULL COMMENT 'Name of estate.',
    user_id        BIGINT UNSIGNED NOT NULL COMMENT 'Id of user.',
    active         TINYINT(1)      NOT NULL DEFAULT 0 COMMENT 'Status of deletion estate.',
    created_at     TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation estate',
    updated_at     TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating estate',

    PRIMARY KEY (id),
    CONSTRAINT rep_estate_types_estate_type_id_name_uk UNIQUE (estate_type_id, name),
    CONSTRAINT rep_estate_types_estate_type_id_name_fk FOREIGN KEY (estate_type_id) REFERENCES real_estate_portfolio.rep_estate_types (id),
    CONSTRAINT rep_estate_types_user_id_name_fk FOREIGN KEY (user_id) REFERENCES real_estate_portfolio.rep_users (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Estate types';