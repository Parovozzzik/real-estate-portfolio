CREATE TABLE IF NOT EXISTS rep_transaction_group_settings
(
    id                BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of transaction group setting.',
    name              VARCHAR(64) NOT NULL COMMENT 'Name of transaction group setting.',
    cost              DECIMAL(16, 2)       DEFAULT NULL COMMENT 'Cost of object.',
    down_payment      DECIMAL(16, 2)       DEFAULT NULL COMMENT 'Down payment.',
    own_funds         DECIMAL(16, 2)       DEFAULT NULL COMMENT 'Own funds.',
    third_party_funds DECIMAL(16, 2)       DEFAULT NULL COMMENT 'Third party funds.',
    interest_rate     DECIMAL(6, 2)        DEFAULT NULL COMMENT 'Interest rate',
    frequency_id      BIGINT UNSIGNED      DEFAULT NULL COMMENT 'Id of frequency.',
    repayment_plan_id BIGINT UNSIGNED      DEFAULT NULL COMMENT 'Id of repayment plan',
    date_start        DATE                 DEFAULT NULL COMMENT 'Date start',
    loan_term         SMALLINT             DEFAULT NULL COMMENT 'Loan term',
    payday            TINYINT(1)           DEFAULT NULL COMMENT 'Day of payment',
    payday_on_workday TINYINT(1)           DEFAULT NULL COMMENT 'Day of payment after holidays',
    created_at        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation transaction group setting',
    updated_at        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating transaction group setting',

    PRIMARY KEY (id),
    CONSTRAINT rep_transaction_groups_settings_frequency_id_fk FOREIGN KEY (frequency_id) REFERENCES real_estate_portfolio.rep_transaction_frequencies (id),
    CONSTRAINT rep_transaction_groups_settings_repayment_plan_id_fk FOREIGN KEY (repayment_plan_id) REFERENCES real_estate_portfolio.rep_transaction_repayment_plans (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'transaction group settings';