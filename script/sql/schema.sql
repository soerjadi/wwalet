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
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS transactions (
    id varchar(55) not null primary key,
    user_id varchar(55) not null,
    status enum('SUCCESS', 'FAILED') not null,
    type enum('DEBIT', 'CREDIT') not null,
    category enum('topup', 'transfer', 'payment') not null,
    amount bigint not null,
    remarks text,
    created_at timestamp default current_timestamp
);