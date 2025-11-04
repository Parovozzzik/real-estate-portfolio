-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS real_estate_portfolio.rep_transactions
(
    id         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of transaction.',
    group_id   BIGINT UNSIGNED NOT NULL COMMENT 'Group id of transaction.',
    type_id    BIGINT UNSIGNED NOT NULL COMMENT 'Type id of transaction.',
    sum        DECIMAL(16, 2)  NOT NULL COMMENT 'Sum of transaction.',
    date       DATE            NOT NULL COMMENT 'Date of transaction.',
    comment    VARCHAR(255)             DEFAULT NULL COMMENT 'Comment of transaction.',
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation transaction.',
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating transaction.',

    PRIMARY KEY (id),
    CONSTRAINT rep_transactions_group_id_fk FOREIGN KEY (group_id) REFERENCES real_estate_portfolio.rep_transaction_groups (id),
    CONSTRAINT rep_transactions_type_id_fk FOREIGN KEY (type_id) REFERENCES real_estate_portfolio.rep_transaction_types (id)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'Transactions';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS real_estate_portfolio.rep_transactions;
-- +goose StatementEnd
