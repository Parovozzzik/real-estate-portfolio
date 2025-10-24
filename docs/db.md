rep_users
- id
- email
- name
- password
- email_confirm
- email_confirm_code
- email_confirmed_at
- deleted
- created_at
- updated_at

==================

rep_estate_types
- id
- name
- active
- created_at
- updated_at

rep_estates
- id
- estate_type_id
- name
- description
- user_d
- active
- created_at
- updated_at

==================

rep_transaction_types (ипотека, рассрочка, кредит, аренда, к/у, фкр, эл-во, ежегодный взнос, разовый взнос, регистрация, налог, возврат налога и т.д.)
- id
- name
- active
- created_at
- updated_at

transaction_frequencies (месяц, квартал, год, 3 года)
- id
- name
- created_at
- updated_at

transaction_group_repayment_plans (аннуитетная, дифференцированная)
- id
- name
- created_at
- updated_at

transaction_group_settings
- id
- name?
- cost
- down_payment
- own_funds
- third_party_funds
- interest_rate
- frequency_id
- repayment_plan_id (аннуитетная, дифференцированная)
- date_start
- payday
- payday_on_workday
- date_end
- created_at
- updated_at

rep_transaction_groups
- id
- estate_id
- transaction_group_setting_id
- direction (1 - доход, 0 - расход)
- regularity (1 - да, 0 - нет)
- created_at
- updated_at

rep_transactions
- id
- transaction_group_id
- transaction_type_id
- sum
- date
- comment
- created_at
- updated_at