CREATE TABLE IF NOT EXISTS rep_transaction_repayment_plans
(
    id         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'Id of transaction repayment plan.',
    name       VARCHAR(64) NOT NULL COMMENT 'Name of transaction repayment plan.',
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date of creation transaction repayment plan',
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Date of updating transaction repayment plan',

    PRIMARY KEY (id),
    CONSTRAINT rep_transaction_repayment_plans_name_uk UNIQUE (name)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT 'Transaction repayment_plans';