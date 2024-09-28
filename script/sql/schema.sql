CREATE TABLE IF NOT EXISTS users (
    id varchar(55) not null primary key,
    first_name varchar(60) not null,
    last_name varchar(60) not null,
    phone_number varchar(13) not null,
    address varchar(155) not null,
    pin varchar(255) not null,
    salt varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS  wallets (
    id varchar(55) not null primary key,
    user_id varchar(55) not null,
    balance bigint not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    constraint fk_wallets_users
        foreign key (user_id)
        references users (id)
);

CREATE TYPE transaction_status as enum ('SUCCESS', 'FAILED');
CREATE TYPE transaction_type as enum ('DEBIT', 'CREDIT');
CREATE TYPE transaction_category as enum ('topup', 'transfer', 'payment');

CREATE TABLE IF NOT EXISTS transactions (
    id varchar(55) not null primary key,
    user_id varchar(55) not null,
    status transaction_status not null,
    type transaction_type not null,
    category transaction_category not null,
    amount bigint not null,
    remarks text,
    created_at timestamp default current_timestamp,
    balance_before bigint not null,
    balance_after bigint not null,
    constraint fk_transactions_users
        foreign key (user_id)
        references users (id)
);