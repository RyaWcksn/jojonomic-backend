CREATE TABLE IF NOT EXISTS tbl_harga (
    reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
    admin_id VARCHAR (15) NOT NULL,
    harga_topup INTEGER NOT NULL,
    harga_buyback INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tbl_rekening (
    norek VARCHAR (15) PRIMARY KEY NOT NULL,
    gold_balance DECIMAL(12, 3) DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(norek)
);

CREATE TABLE IF NOT EXISTS tbl_transaksi (
    reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
    norek VARCHAR (15) NOT NULL,
    type VARCHAR (15) NOT NULL,
    harga_topup INTEGER NOT NULL,
    harga_buyback INTEGER NOT NULL,
    gold_weight DECIMAL(12, 3) DEFAULT 0,
    gold_balance DECIMAL(12, 3) DEFAULT 0,
    created_at INTEGER NOT NULL
);

INSERT INTO tbl_rekening (norek, gold_balance, created_at) VALUES ( 'rek001', 0.001, NOW());
