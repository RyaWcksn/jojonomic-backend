CREATE TABLE IF NOT EXISTS tbl_harga (
    reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
    admin_id VARCHAR (15) NOT NULL,
    harga_topup DECIMAL(12, 2) NOT NULL,
    harga_buyback DECIMAL(12, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tbl_rekening (
    norek VARCHAR (15) PRIMARY UNIQUE NOT NULL,
    gold_balance DECIMAL(12, 2) DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tbl_transaksi (
    reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
    norek VARCHAR (15) NOT NULL,
    type VARCHAR (15) NOT NULL,
    gold_weight DECIMAL(12, 2) DEFAULT 0,
    gold_balance DECIMAL(12, 2) DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)

