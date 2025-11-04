-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS real_estate_portfolio.rep_transaction_groups
(
    id         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of transaction group setting.',
    estate_id  BIGINT UNSIGNED NOT NULL COMMENT 'Name of transaction group setting.',
    setting_id BIGINT UNSIGNED          DEFAULT NULL COMMENT 'Cost of object.',
    direction  TINYINT(1)      NOT NULL DEFAULT 0 COMMENT 'Direction of transaction type (1 - in, 0 - out)',
    regularity TINYINT(1)      NOT NULL DEFAULT 0 COMMENT 'Regularity of transaction type (1 - regular, 0 - one-time)',
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation transaction group setting',
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating transaction group setting',

    PRIMARY KEY (id),
    CONSTRAINT rep_transaction_groups_estate_id_name_fk FOREIGN KEY (estate_id) REFERENCES real_estate_portfolio.rep_estates (id),
    CONSTRAINT rep_transaction_groups_setting_id_name_fk FOREIGN KEY (setting_id) REFERENCES real_estate_portfolio.rep_transaction_group_settings (id)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'Transaction groups';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS real_estate_portfolio.rep_transaction_groups;
-- +goose StatementEnd
