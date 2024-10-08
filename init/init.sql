CREATE TABLE IF NOT EXISTS "order" (
    order_uid VARCHAR not NULL,
    track_number VARCHAR not NULL,
    entry VARCHAR not NULL,
    delivery_id BIGINT not NULL,
    payment_id BIGINT not NULL,
    locale VARCHAR(2) not NULL,
    internal_signature VARCHAR,
    customer_id VARCHAR,
    delivery_service VARCHAR,
    shardkey VARCHAR,
    sm_id int,
    date_created timestamptz DEFAULT CURRENT_TIMESTAMP not NULL,
    oof_shard VARCHAR
);

CREATE TABLE IF NOT EXISTS item (
    id BIGINT not NULL,
    order_uid VARCHAR not NULL,
    chrt_id BIGINT not NULL,
    track_number VARCHAR not NULL,
    price BIGINT,
    rid VARCHAR,
    sale INT,
    size VARCHAR(3),
    total_price BIGINT,
    nm_id BIGINT,
    brand VARCHAR(50),
    status VARCHAR(3)
);

CREATE TABLE IF NOT EXISTS payment (
    id BIGINT not NULL,
    transaction VARCHAR not NULL,
    request_id VARCHAR,
    currency VARCHAR(3) not NULL,
    provider VARCHAR not NULL,
    amount BIGINT,
    payment_dt BIGINT,
    bank VARCHAR(30),
    delivery_cost INT,
    goods_total BIGINT,
    custom_fee INT
);

CREATE TABLE IF NOT EXISTS delivery (
    id BIGINT not NULL,
    name VARCHAR not NULL,
    phone VARCHAR(20) not NULL,
    zip VARCHAR(12) not NULL,
    city VARCHAR(40) not NULL,
    address VARCHAR(60) not NULL,
    region VARCHAR(40) not NULL,
    email VARCHAR not NULL
);
