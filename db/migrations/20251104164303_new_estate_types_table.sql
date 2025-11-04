-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS real_estate_portfolio.rep_estate_types
(
    id         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of estate type.',
    name       VARCHAR(64) NOT NULL COMMENT 'Name of estate type.',
    icon       VARCHAR(32) NOT NULL COMMENT 'Code of icon.',
    active     TINYINT(1)  NOT NULL DEFAULT 0 COMMENT 'Status of deletion estate type.',
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation estate type',
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating estate type',

    PRIMARY KEY (id),
    CONSTRAINT rep_estate_types_name_uk UNIQUE (name)
    ) ENGINE = INNODB
    DEFAULT CHARSET = utf8 COMMENT 'Estate types';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS real_estate_portfolio.rep_estate_types;

-- +goose StatementEnd
