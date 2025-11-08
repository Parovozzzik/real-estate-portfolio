-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS real_estate_portfolio.rep_estate_values
(
    id                 BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of estate value.',
    estate_id          BIGINT UNSIGNED NOT NULL COMMENT 'Id of estate.',
    date               DATE            NOT NULL COMMENT 'Date of estate values.',
    income             DECIMAL(16, 2)  NOT NULL COMMENT 'Income in this month.',
    expense            DECIMAL(16, 2)  NOT NULL COMMENT 'Expense in this month.',
    profit             DECIMAL(16, 2)  NOT NULL COMMENT 'Profit in this month.',
    cumulative_income  DECIMAL(16, 2)  NOT NULL COMMENT 'Cumulative income.',
    cumulative_expense DECIMAL(16, 2)  NOT NULL COMMENT 'Cumulative expense.',
    cumulative_profit  DECIMAL(16, 2)  NOT NULL COMMENT 'Cumulative profit.',
    roi                DECIMAL(7, 2)   NOT NULL COMMENT 'Return on investment.',

    created_at         TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation estate value.',
    updated_at         TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating estate value.',

    PRIMARY KEY (id),
    CONSTRAINT rep_estate_estate_id_date_uk UNIQUE (estate_id, date),
    CONSTRAINT rep_estate_estate_id_fk FOREIGN KEY (estate_id) REFERENCES real_estate_portfolio.rep_estates (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Estate values';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS real_estate_portfolio.rep_estate_values;
-- +goose StatementEnd
